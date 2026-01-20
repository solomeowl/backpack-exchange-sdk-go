// Package auth provides ED25519 signing functionality for Backpack Exchange API.
package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Signer handles ED25519 signing for API requests.
type Signer struct {
	publicKey  string
	privateKey ed25519.PrivateKey
}

// NewSigner creates a new Signer from base64-encoded keys.
func NewSigner(publicKey, secretKey string) (*Signer, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode secret key: %w", err)
	}

	// ED25519 seed is 32 bytes, full private key is 64 bytes
	var privateKey ed25519.PrivateKey
	if len(privateKeyBytes) == ed25519.SeedSize {
		privateKey = ed25519.NewKeyFromSeed(privateKeyBytes)
	} else if len(privateKeyBytes) == ed25519.PrivateKeySize {
		privateKey = privateKeyBytes
	} else {
		return nil, fmt.Errorf("invalid secret key length: expected %d or %d bytes, got %d",
			ed25519.SeedSize, ed25519.PrivateKeySize, len(privateKeyBytes))
	}

	return &Signer{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

// PublicKey returns the public key (API key).
func (s *Signer) PublicKey() string {
	return s.publicKey
}

// Sign signs a message and returns a base64-encoded signature.
func (s *Signer) Sign(message string) string {
	signature := ed25519.Sign(s.privateKey, []byte(message))
	return base64.StdEncoding.EncodeToString(signature)
}

// BuildSigningString builds the signing string for an API request.
// Parameters are sorted alphabetically and boolean values are converted to lowercase strings.
func BuildSigningString(instruction string, params map[string]any, timestamp, window int64) string {
	var sb strings.Builder

	sb.WriteString("instruction=")
	sb.WriteString(instruction)

	if len(params) > 0 {
		// Sort keys alphabetically
		keys := make([]string, 0, len(params))
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			sb.WriteString("&")
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(formatValue(params[k]))
		}
	}

	sb.WriteString("&timestamp=")
	sb.WriteString(strconv.FormatInt(timestamp, 10))
	sb.WriteString("&window=")
	sb.WriteString(strconv.FormatInt(window, 10))

	return sb.String()
}

// BuildBatchSigningString builds the signing string for batch order execution.
// Each order gets its own instruction prefix and they are concatenated together.
func BuildBatchSigningString(orders []map[string]any, timestamp, window int64) string {
	var sb strings.Builder

	for i, order := range orders {
		if i > 0 {
			sb.WriteString("&")
		}

		sb.WriteString("instruction=orderExecute")

		// Sort keys alphabetically
		keys := make([]string, 0, len(order))
		for k := range order {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			sb.WriteString("&")
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(formatValue(order[k]))
		}
	}

	sb.WriteString("&timestamp=")
	sb.WriteString(strconv.FormatInt(timestamp, 10))
	sb.WriteString("&window=")
	sb.WriteString(strconv.FormatInt(window, 10))

	return sb.String()
}

// formatValue formats a value for signing.
// Boolean values are converted to lowercase strings ("true"/"false").
func formatValue(v any) string {
	switch val := v.(type) {
	case bool:
		if val {
			return "true"
		}
		return "false"
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", val)
	}
}
