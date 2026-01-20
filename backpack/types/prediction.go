package types

// PredictionEvent represents a prediction event from GET /api/v1/prediction.
type PredictionEvent struct {
	Slug              string              `json:"slug"`
	Title             string              `json:"title"`
	PredictionMarkets []PredictionMarket  `json:"predictionMarkets"`
	Tags              []PredictionTag     `json:"tags"`
	Series            []PredictionSeries  `json:"series"`
	Description       string              `json:"description"`
	ImgURL            string              `json:"imgUrl,omitempty"`
	EstimatedEndDate  string              `json:"estimatedEndDate,omitempty"`
	QuoteVolume       string              `json:"quoteVolume"`
	Resolved          bool                `json:"resolved"`
}

// PredictionMarket represents an individual prediction market within an event.
type PredictionMarket struct {
	MarketSymbol    string `json:"marketSymbol"`
	Question        string `json:"question"`
	GroupLabel      string `json:"groupLabel,omitempty"`
	YesOutcomeLabel string `json:"yesOutcomeLabel"`
	NoOutcomeLabel  string `json:"noOutcomeLabel"`
	Rules           string `json:"rules"`
	ResolvedAt      string `json:"resolvedAt,omitempty"`
	ResolutionPrice string `json:"resolutionPrice,omitempty"`
	ActivePrice     string `json:"activePrice"`
	QuoteVolume     string `json:"quoteVolume"`
	ImgURL          string `json:"imgUrl,omitempty"`
}

// PredictionTag represents a tag associated with a prediction event.
type PredictionTag struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// PredictionSeries represents a series that a prediction event belongs to.
type PredictionSeries struct {
	Slug       string `json:"slug,omitempty"`
	Title      string `json:"title,omitempty"`
	Recurrence string `json:"recurrence,omitempty"`
}
