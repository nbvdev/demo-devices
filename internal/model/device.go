package model

import "time"

type Device struct {
	Id      int64     `json:"id,omitempty"`
	Name    string    `json:"name"`
	Brand   string    `json:"brand"`
	Created time.Time `json:"created"`
}

func (d *Device) IsSuitableForUpdate() bool {
	return d != nil && d.Brand != "" && d.Name != ""
}

func (d *Device) Patch(device *Device) {
	if device.Name != "" {
		d.Name = device.Name
	}
	if device.Brand != "" {
		d.Brand = device.Brand
	}
}
