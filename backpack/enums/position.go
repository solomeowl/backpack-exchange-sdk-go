package enums

// PositionState represents the state of a position (for history).
type PositionState string

const (
	PositionStateOpen   PositionState = "Open"
	PositionStateClosed PositionState = "Closed"
)
