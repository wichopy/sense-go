// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    device, err := UnmarshalDevice(bytes)
//    bytes, err = device.Marshal()

package app

import "encoding/json"

type DeviceResponse []DeviceElement

func UnmarshalDevices(data []byte) (DeviceResponse, error) {
	var r DeviceResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeviceResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeviceElement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Tags Tags   `json:"tags"`
}
