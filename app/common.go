package app

type Tags struct {
	Alertable                   *string     `json:"Alertable,omitempty"`
	AlwaysOn                    *string     `json:"AlwaysOn,omitempty"`
	DateCreated                 *string     `json:"DateCreated,omitempty"`
	DateFirstUsage              *string     `json:"DateFirstUsage,omitempty"`
	DefaultLocation             interface{} `json:"DefaultLocation"`
	DefaultMake                 interface{} `json:"DefaultMake"`
	DefaultModel                interface{} `json:"DefaultModel"`
	DefaultUserDeviceType       string      `json:"DefaultUserDeviceType"`
	DeployToMonitor             *string     `json:"DeployToMonitor,omitempty"`
	DeviceListAllowed           string      `json:"DeviceListAllowed"`
	Mature                      *string     `json:"Mature,omitempty"`
	ModelCreatedVersion         *string     `json:"ModelCreatedVersion,omitempty"`
	ModelUpdatedVersion         *string     `json:"ModelUpdatedVersion,omitempty"`
	NameUseredit                *string     `json:"name_useredit,omitempty"`
	OriginalName                *string     `json:"OriginalName,omitempty"`
	PeerNames                   []PeerName  `json:"PeerNames"`
	Pending                     *string     `json:"Pending,omitempty"`
	PreselectionIndex           *int64      `json:"PreselectionIndex,omitempty"`
	Revoked                     *string     `json:"Revoked,omitempty"`
	TimelineAllowed             string      `json:"TimelineAllowed"`
	TimelineDefault             *string     `json:"TimelineDefault,omitempty"`
	Type                        *string     `json:"Type,omitempty"`
	UserDeletable               *string     `json:"UserDeletable,omitempty"`
	UserDeviceType              string      `json:"UserDeviceType"`
	UserDeviceTypeDisplayString string      `json:"UserDeviceTypeDisplayString"`
	UserEditable                string      `json:"UserEditable"`
	UserEditableMeta            *string     `json:"UserEditableMeta,omitempty"`
	UserMergeable               *string     `json:"UserMergeable,omitempty"`
	NameUserGuess               *string     `json:"NameUserGuess,omitempty"`
}

type PeerName struct {
	Name                        string `json:"Name"`
	UserDeviceType              string `json:"UserDeviceType"`
	Percent                     float64  `json:"Percent"`
	Icon                        string `json:"Icon"`
	UserDeviceTypeDisplayString string `json:"UserDeviceTypeDisplayString"`
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
