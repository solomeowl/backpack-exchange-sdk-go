package types

// PredictionMarket represents a prediction market group from the API.
type PredictionMarket struct {
	Description       string                  `json:"description,omitempty"`
	EstimatedEndDate  string                  `json:"estimatedEndDate,omitempty"`
	ImgURL            *string                 `json:"imgUrl,omitempty"`
	PredictionMarkets []PredictionMarketItem  `json:"predictionMarkets,omitempty"`
	QuoteVolume       string                  `json:"quoteVolume,omitempty"`
	Resolved          bool                    `json:"resolved,omitempty"`
	Series            []PredictionSeries      `json:"series,omitempty"`
	Slug              string                  `json:"slug,omitempty"`
	Tags              []PredictionMarketTag   `json:"tags,omitempty"`
	Title             string                  `json:"title,omitempty"`
}

// PredictionMarketItem represents an individual prediction market within a group.
type PredictionMarketItem struct {
	ActivePrice     string  `json:"activePrice,omitempty"`
	GroupLabel      string  `json:"groupLabel,omitempty"`
	ImgURL          *string `json:"imgUrl,omitempty"`
	MarketSymbol    string  `json:"marketSymbol,omitempty"`
	NoOutcomeLabel  string  `json:"noOutcomeLabel,omitempty"`
	Question        string  `json:"question,omitempty"`
	QuoteVolume     string  `json:"quoteVolume,omitempty"`
	ResolutionPrice string  `json:"resolutionPrice,omitempty"`
	ResolvedAt      string  `json:"resolvedAt,omitempty"`
	Rules           string  `json:"rules,omitempty"`
	YesOutcomeLabel string  `json:"yesOutcomeLabel,omitempty"`
}

// PredictionMarketTag represents a tag associated with a prediction market group.
type PredictionMarketTag struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// PredictionSeries represents a series that a prediction market belongs to.
type PredictionSeries struct {
	Recurrence string `json:"recurrence,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Title      string `json:"title,omitempty"`
}

// PredictionTag represents a prediction market tag from the tags endpoint.
type PredictionTag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}
