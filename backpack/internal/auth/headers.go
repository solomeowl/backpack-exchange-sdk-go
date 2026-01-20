package auth

import (
	"strconv"
	"time"
)

// Headers represents authentication headers for API requests.
type Headers struct {
	APIKey    string
	Signature string
	Timestamp string
	Window    string
}

// ToMap converts Headers to a map for HTTP headers.
func (h Headers) ToMap() map[string]string {
	return map[string]string{
		"X-API-Key":    h.APIKey,
		"X-Signature":  h.Signature,
		"X-Timestamp":  h.Timestamp,
		"X-Window":     h.Window,
		"Content-Type": "application/json; charset=utf-8",
	}
}

// GenerateHeaders generates authentication headers for a standard API request.
func (s *Signer) GenerateHeaders(instruction string, params map[string]any, window int64) Headers {
	timestamp := time.Now().UnixMilli()
	signStr := BuildSigningString(instruction, params, timestamp, window)
	signature := s.Sign(signStr)

	return Headers{
		APIKey:    s.publicKey,
		Signature: signature,
		Timestamp: strconv.FormatInt(timestamp, 10),
		Window:    strconv.FormatInt(window, 10),
	}
}

// GenerateBatchHeaders generates authentication headers for batch order execution.
func (s *Signer) GenerateBatchHeaders(orders []map[string]any, window int64) Headers {
	timestamp := time.Now().UnixMilli()
	signStr := BuildBatchSigningString(orders, timestamp, window)
	signature := s.Sign(signStr)

	return Headers{
		APIKey:    s.publicKey,
		Signature: signature,
		Timestamp: strconv.FormatInt(timestamp, 10),
		Window:    strconv.FormatInt(window, 10),
	}
}

// GenerateWSSignature generates a signature for WebSocket authentication.
func (s *Signer) GenerateWSSignature(window int64) (signature string, timestamp int64) {
	timestamp = time.Now().UnixMilli()
	signStr := BuildSigningString("subscribe", nil, timestamp, window)
	signature = s.Sign(signStr)
	return signature, timestamp
}
