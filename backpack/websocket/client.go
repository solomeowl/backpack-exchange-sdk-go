// Package websocket provides WebSocket client for Backpack Exchange real-time data.
package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/internal/auth"
)

const (
	DefaultWSURL       = "wss://ws.backpack.exchange"
	DefaultWindow      = int64(5000)
	pingInterval       = 30 * time.Second
	pongWait           = 120 * time.Second
	writeWait          = 10 * time.Second
	reconnectMinDelay  = 1 * time.Second
	reconnectMaxDelay  = 30 * time.Second
	serverShutdownWait = 30 * time.Second
)

// Client is a WebSocket client for Backpack Exchange.
type Client struct {
	url       string
	signer    *auth.Signer
	window    int64
	conn      *websocket.Conn
	mu        sync.RWMutex
	callbacks map[string][]func(json.RawMessage)
	privateStreams map[string]bool
	done      chan struct{}
	reconnect bool
	connected bool
	dialer    *websocket.Dialer
}

// Option is a functional option for configuring the WebSocket client.
type Option func(*Client) error

// WithWSURL sets a custom WebSocket URL.
func WithWSURL(url string) Option {
	return func(c *Client) error {
		c.url = url
		return nil
	}
}

// WithCredentials sets API credentials for authenticated streams.
func WithCredentials(publicKey, secretKey string) Option {
	return func(c *Client) error {
		signer, err := auth.NewSigner(publicKey, secretKey)
		if err != nil {
			return err
		}
		c.signer = signer
		return nil
	}
}

// WithWSWindow sets the signature validity window.
func WithWSWindow(window int64) Option {
	return func(c *Client) error {
		c.window = window
		return nil
	}
}

// WithAutoReconnect enables/disables automatic reconnection.
func WithAutoReconnect(enabled bool) Option {
	return func(c *Client) error {
		c.reconnect = enabled
		return nil
	}
}

// NewClient creates a new WebSocket client.
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		url:       DefaultWSURL,
		window:    DefaultWindow,
		callbacks: make(map[string][]func(json.RawMessage)),
		privateStreams: make(map[string]bool),
		done:      make(chan struct{}),
		reconnect: true,
		dialer:    websocket.DefaultDialer,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Connect establishes a WebSocket connection.
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	conn, resp, err := c.dialer.DialContext(ctx, c.url, http.Header{})
	if err != nil {
		if resp != nil {
			return fmt.Errorf("websocket dial failed with status %d: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("websocket dial failed: %w", err)
	}

	c.conn = conn
	c.connected = true
	c.done = make(chan struct{})

	// Start message reader
	go c.readLoop()

	// Start ping sender
	go c.pingLoop()

	return nil
}

// Close closes the WebSocket connection.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.reconnect = false

	if c.conn != nil {
		close(c.done)
		err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			return err
		}
		c.connected = false
		return c.conn.Close()
	}

	return nil
}

// Subscribe subscribes to one or more streams.
func (c *Client) Subscribe(streams []string, callback func(json.RawMessage)) error {
	return c.subscribe(streams, callback, false)
}

// SubscribePrivate subscribes to private authenticated streams.
func (c *Client) SubscribePrivate(streams []string, callback func(json.RawMessage)) error {
	if c.signer == nil {
		return fmt.Errorf("signer required for private streams")
	}
	return c.subscribe(streams, callback, true)
}

func (c *Client) subscribe(streams []string, callback func(json.RawMessage), private bool) error {
	c.mu.Lock()

	// Register callbacks
	for _, stream := range streams {
		c.callbacks[stream] = append(c.callbacks[stream], callback)
		if private {
			c.privateStreams[stream] = true
		} else if _, ok := c.privateStreams[stream]; !ok {
			c.privateStreams[stream] = false
		}
	}

	// Build subscription message
	msg := map[string]any{
		"method": "SUBSCRIBE",
		"params": streams,
	}

	// Add authentication for private streams
	if private && c.signer != nil {
		signature, timestamp := c.signer.GenerateWSSignature(c.window)
		msg["signature"] = []string{
			c.signer.PublicKey(),
			signature,
			strconv.FormatInt(timestamp, 10),
			strconv.FormatInt(c.window, 10),
		}
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.mu.Unlock()
		return err
	}

	conn := c.conn
	c.mu.Unlock()

	if conn == nil {
		return fmt.Errorf("not connected")
	}

	return conn.WriteMessage(websocket.TextMessage, msgBytes)
}

// Unsubscribe unsubscribes from one or more streams.
func (c *Client) Unsubscribe(streams []string) error {
	c.mu.Lock()

	// Remove callbacks
	for _, stream := range streams {
		delete(c.callbacks, stream)
		delete(c.privateStreams, stream)
	}

	msg := map[string]any{
		"method": "UNSUBSCRIBE",
		"params": streams,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.mu.Unlock()
		return err
	}

	conn := c.conn
	c.mu.Unlock()

	if conn == nil {
		return fmt.Errorf("not connected")
	}

	return conn.WriteMessage(websocket.TextMessage, msgBytes)
}

func (c *Client) readLoop() {
	defer func() {
		c.mu.Lock()
		c.connected = false
		c.mu.Unlock()

		if c.reconnect {
			go c.reconnectLoop()
		}
	}()

	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		select {
		case <-c.done:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					return
				}
				return
			}

			c.handleMessage(message)
		}
	}
}

func (c *Client) handleMessage(message []byte) {
	var msg struct {
		Stream string          `json:"stream"`
		Data   json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(message, &msg); err != nil {
		return
	}

	if msg.Stream == "" {
		return
	}

	c.mu.RLock()
	callbacks := c.callbacks[msg.Stream]
	c.mu.RUnlock()

	for _, cb := range callbacks {
		go cb(msg.Data)
	}
}

func (c *Client) pingLoop() {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.done:
			return
		case <-ticker.C:
			c.mu.RLock()
			conn := c.conn
			c.mu.RUnlock()

			if conn != nil {
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait)); err != nil {
					return
				}
			}
		}
	}
}

func (c *Client) reconnectLoop() {
	delay := reconnectMinDelay

	for {
		c.mu.RLock()
		shouldReconnect := c.reconnect
		c.mu.RUnlock()

		if !shouldReconnect {
			return
		}

		time.Sleep(delay)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := c.Connect(ctx)
		cancel()

		if err == nil {
			// Resubscribe to all streams
			c.mu.RLock()
			publicStreams := make([]string, 0, len(c.callbacks))
			privateStreams := make([]string, 0, len(c.callbacks))
			for stream := range c.callbacks {
				if c.privateStreams[stream] {
					privateStreams = append(privateStreams, stream)
				} else {
					publicStreams = append(publicStreams, stream)
				}
			}
			c.mu.RUnlock()

			if len(publicStreams) > 0 {
				msg := map[string]any{
					"method": "SUBSCRIBE",
					"params": publicStreams,
				}
				msgBytes, _ := json.Marshal(msg)
				c.conn.WriteMessage(websocket.TextMessage, msgBytes)
			}
			if len(privateStreams) > 0 && c.signer != nil {
				signature, timestamp := c.signer.GenerateWSSignature(c.window)
				msg := map[string]any{
					"method": "SUBSCRIBE",
					"params": privateStreams,
					"signature": []string{
						c.signer.PublicKey(),
						signature,
						strconv.FormatInt(timestamp, 10),
						strconv.FormatInt(c.window, 10),
					},
				}
				msgBytes, _ := json.Marshal(msg)
				c.conn.WriteMessage(websocket.TextMessage, msgBytes)
			}

			return
		}

		// Exponential backoff
		delay *= 2
		if delay > reconnectMaxDelay {
			delay = reconnectMaxDelay
		}
	}
}

// IsConnected returns whether the client is connected.
func (c *Client) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connected
}
