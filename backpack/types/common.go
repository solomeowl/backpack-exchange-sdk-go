// Package types provides type definitions for Backpack Exchange SDK.
package types

// Pagination represents pagination parameters.
type Pagination struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// TimeRange represents a time range for queries.
type TimeRange struct {
	From int64 `json:"from,omitempty"` // Unix timestamp in milliseconds
	To   int64 `json:"to,omitempty"`   // Unix timestamp in milliseconds
}

// PaginatedResponse represents a paginated response.
type PaginatedResponse[T any] struct {
	Data       []T `json:"data"`
	TotalCount int `json:"totalCount,omitempty"`
}
