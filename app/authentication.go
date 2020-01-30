package app

import "encoding/json"

func UnmarshalSenseAuthenticationResponse(data []byte) (SenseAuthenticationResponse, error) {
	var r SenseAuthenticationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SenseAuthenticationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Auth struct {
	AccessToken string `json:"accessToken"`
	AccountID   int64  `json:"accountID"`
	UserID      int64  `json:"userID"`
	MonitorID   int64  `json:"monitorID"`
}

type SenseAuthenticationResponse struct {
	Authorized   bool        `json:"authorized"`
	AccountID    int64       `json:"account_id"`
	UserID       int64       `json:"user_id"`
	AccessToken  string      `json:"access_token"`
	Monitors     []Monitor   `json:"monitors"`
	BridgeServer string      `json:"bridge_server"`
	PartnerID    interface{} `json:"partner_id"`
	DateCreated  string      `json:"date_created"`
	AbCohort     string      `json:"ab_cohort"`
}

type Monitor struct {
	ID                int64         `json:"id"`
	SerialNumber      string        `json:"serial_number"`
	TimeZone          string        `json:"time_zone"`
	SolarConnected    bool          `json:"solar_connected"`
	SolarConfigured   bool          `json:"solar_configured"`
	Online            bool          `json:"online"`
	Attributes        Attributes    `json:"attributes"`
	DataSharing       []interface{} `json:"data_sharing"`
	EthernetSupported bool          `json:"ethernet_supported"`
}

type Attributes struct {
	ID                interface{} `json:"id"`
	Name              interface{} `json:"name"`
	State             interface{} `json:"state"`
	Cost              float64     `json:"cost"`
	UserSetCost       bool        `json:"user_set_cost"`
	CycleStart        interface{} `json:"cycle_start"`
	BasementType      interface{} `json:"basement_type"`
	HomeSizeType      interface{} `json:"home_size_type"`
	HomeType          interface{} `json:"home_type"`
	NumberOfOccupants interface{} `json:"number_of_occupants"`
	OccupancyType     interface{} `json:"occupancy_type"`
	YearBuiltType     interface{} `json:"year_built_type"`
	PostalCode        interface{} `json:"postal_code"`
	ElectricityCost   interface{} `json:"electricity_cost"`
}

func UnmarshalSenseAuthenticationRequest(data []byte) (SenseAuthenticationRequest, error) {
	var r SenseAuthenticationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SenseAuthenticationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SenseAuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
