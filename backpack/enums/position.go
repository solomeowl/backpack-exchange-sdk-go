package enums

// PositionSide represents the side of a position.
type PositionSide string

const (
	PositionSideLong  PositionSide = "Long"
	PositionSideShort PositionSide = "Short"
)

// MarginType represents the type of margin.
type MarginType string

const (
	MarginTypeCross    MarginType = "Cross"
	MarginTypeIsolated MarginType = "Isolated"
)

// PositionState represents the state of a position (for history).
type PositionState string

const (
	PositionStateOpen   PositionState = "Open"
	PositionStateClosed PositionState = "Closed"
)
