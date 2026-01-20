package types

// PredictionMarket represents a prediction market.
type PredictionMarket struct {
	Symbol        string   `json:"symbol"`
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	ResolutionTime int64   `json:"resolutionTime,omitempty"`
	Status        string   `json:"status,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

// PredictionTag represents a prediction market tag.
type PredictionTag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}
