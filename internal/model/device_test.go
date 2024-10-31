package model

import (
	"testing"
)

func TestDevice_IsSuitableForUpdate(t *testing.T) {
	patterns := []struct {
		TestDevice Device
		IsSuitable bool
	}{
		{
			Device{}, false,
		},
		{
			Device{Name: "name"}, false,
		},
		{
			Device{Brand: "brand"}, false,
		},
		{
			Device{Name: "name", Brand: "brand"}, true,
		},
	}
	for _, p := range patterns {
		if p.IsSuitable != p.TestDevice.IsSuitableForUpdate() {
			t.Error("expected ", p.IsSuitable, "got ", p.TestDevice.IsSuitableForUpdate(), "entity ", p.TestDevice)
		}
	}
}

func TestDevice_Merge(t *testing.T) {
	patterns := []struct {
		SourceDevice   Device
		PatchDevice    Device
		ExpectedDevice Device
	}{
		{
			Device{}, Device{}, Device{},
		},
		{
			Device{Name: "n1", Brand: "b1"},
			Device{},
			Device{Name: "n1", Brand: "b1"},
		},
		{
			Device{Name: "n1", Brand: "b1"},
			Device{Name: "n2", Brand: "b2"},
			Device{Name: "n2", Brand: "b2"},
		},
	}

	for _, p := range patterns {
		merged := p.SourceDevice
		merged.Patch(&p.PatchDevice)
		if merged != p.ExpectedDevice {
			t.Error("expected ", p.ExpectedDevice, "got ", merged)
		}
	}
}
