package types

type Tags struct {
	DefaultUserDeviceType       string `json:"DefaultUserDeviceType"`
	DeviceListAllowed           string `json:"DeviceListAllowed"`
	TimelineAllowed             string `json:"TimelineAllowed"`
	UserDeviceType              string `json:"UserDeviceType"`
	UserDeviceTypeDisplayString string `json:"UserDeviceTypeDisplayString"`
	UserEditable                string `json:"UserEditable"`
}

type Device struct {
	ID        string        `json:"id"`
	MonitorID int64         `json:"monitorId"`
	Name      string        `json:"name"`
	Icon      string        `json:"icon"`
	Tags      Tags          `json:"tags"`
	History   []int64       `json:"history"`
	Avgw      int64         `json:"avgw"`
	Attrs     []interface{} `json:"attrs"`
	W         float64       `json:"w"`
}
