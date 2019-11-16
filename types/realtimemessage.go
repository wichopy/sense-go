package types

import "encoding/json"

type RealTimeMessageType string

const UpdateMessageTime = RealTimeMessageType("realtime_update")

func UnmarshalSenseRealTimeMessage(data []byte) (SenseRealTimeMessage, error) {
	var r SenseRealTimeMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SenseRealTimeMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SenseRealTimeMessage struct {
	Payload Payload             `json:"payload"`
	Type    RealTimeMessageType `json:"type"`
}

type Payload struct {
	Voltage  []float64          `json:"voltage"`
	Frame    int64              `json:"frame"`
	Devices  []Device           `json:"devices"`
	Deltas   []interface{}      `json:"deltas"`
	Channels []float64          `json:"channels"`
	Hz       float64            `json:"hz"`
	W        float64            `json:"w"`
	Stats    map[string]float64 `json:"_stats"`
	Epoch    int64              `json:"epoch"`
}

type Device struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Icon  string        `json:"icon"`
	Tags  Tags          `json:"tags"`
	Attrs []interface{} `json:"attrs"`
	W     float64       `json:"w"`
}

type Tags struct {
	DefaultUserDeviceType       string `json:"DefaultUserDeviceType"`
	DeviceListAllowed           string `json:"DeviceListAllowed"`
	TimelineAllowed             string `json:"TimelineAllowed"`
	UserDeviceType              string `json:"UserDeviceType"`
	UserDeviceTypeDisplayString string `json:"UserDeviceTypeDisplayString"`
	UserEditable                string `json:"UserEditable"`
}
