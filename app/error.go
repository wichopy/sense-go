package app

import "encoding/json"

func UnmarshalSenseError(data []byte) (SenseError, error) {
	var r SenseError
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SenseError) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SenseError struct {
	Status      string `json:"status"`
	ErrorReason string `json:"error_reason"`
}
