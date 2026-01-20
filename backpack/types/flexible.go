package types

import "encoding/json"

// TimeString accepts either a JSON string or number and stores it as a string.
type TimeString string

// UnmarshalJSON allows TimeString to handle both numeric and string timestamps.
func (t *TimeString) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*t = ""
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*t = TimeString(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	*t = TimeString(n.String())
	return nil
}
