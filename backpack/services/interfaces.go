// Package services provides service implementations for Backpack Exchange SDK.
package services

import "context"

// HTTPClient defines the interface for HTTP operations.
type HTTPClient interface {
	Get(ctx context.Context, path string, params map[string]string, result any) error
	GetAuthenticated(ctx context.Context, path string, params map[string]string, instruction string, result any) error
	Post(ctx context.Context, path string, body any, result any) error
	PostAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error
	Delete(ctx context.Context, path string, body any, result any) error
	DeleteAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error
	Patch(ctx context.Context, path string, body any, result any) error
	PatchAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error
	PostBatchOrders(ctx context.Context, path string, orders []map[string]any, result any) error
}
