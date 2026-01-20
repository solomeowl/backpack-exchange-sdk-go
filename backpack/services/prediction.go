package services

import (
	"context"
	"strconv"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// PredictionService provides prediction market operations (public).
type PredictionService struct {
	client HTTPClient
}

// NewPredictionService creates a new PredictionService.
func NewPredictionService(client HTTPClient) *PredictionService {
	return &PredictionService{client: client}
}

// GetEventsParams represents parameters for getting prediction events.
type GetEventsParams struct {
	Symbol     string // Optional: market symbol
	TagSlug    string // Optional: tag slug
	EventSlug  string // Optional: event slug
	SeriesSlug string // Optional: series slug
	Resolved   *bool  // Optional: filter by resolved status
	Limit      int    // Optional: default 100, max 1000
	Offset     int    // Optional: default 0
}

// GetEvents retrieves prediction events.
func (s *PredictionService) GetEvents(ctx context.Context, params *GetEventsParams) ([]types.PredictionEvent, error) {
	var result []types.PredictionEvent
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.TagSlug != "" {
			queryParams["tagSlug"] = params.TagSlug
		}
		if params.EventSlug != "" {
			queryParams["eventSlug"] = params.EventSlug
		}
		if params.SeriesSlug != "" {
			queryParams["seriesSlug"] = params.SeriesSlug
		}
		if params.Resolved != nil {
			queryParams["resolved"] = strconv.FormatBool(*params.Resolved)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
	}
	if err := s.client.Get(ctx, "api/v1/prediction", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTagsParams represents parameters for getting prediction tags.
type GetTagsParams struct {
	Limit  int // Optional: default 100, max 1000
	Offset int // Optional: default 0
}

// GetTags retrieves prediction market tags.
func (s *PredictionService) GetTags(ctx context.Context, params *GetTagsParams) ([]types.PredictionTag, error) {
	var result []types.PredictionTag
	queryParams := make(map[string]string)
	if params != nil {
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
	}
	if err := s.client.Get(ctx, "api/v1/prediction/tags", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}
