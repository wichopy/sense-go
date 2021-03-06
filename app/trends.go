package app

import "encoding/json"

func UnmarshalTrends(data []byte) (Trends, error) {
	var r Trends
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Trends) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Trends struct {
	Steps       int64       `json:"steps"`
	Start       string      `json:"start"`
	End         string      `json:"end"`
	Consumption Consumption `json:"consumption"`
	Production  Production  `json:"production"`
	TrendText   interface{} `json:"trend_text"`
	UsageText   interface{} `json:"usage_text"`
	Scale       string      `json:"scale"`
}

type Consumption struct {
	Total   int64    `json:"total"`
	Totals  []int64  `json:"totals"`
	Devices []Device `json:"devices"`
}

type Production struct {
	Total int64 `json:"total"`
}
