package service

import (
	"devices/internal/model"
	"devices/mocks/devices/interfaces"
	"errors"
	"testing"
	"time"
)

func TestDeviceService_GetAll(t *testing.T) {
	repo := interfaces.MockDeviceRepository{}
	srv := NewDeviceService(&repo)

	repo.On("List", GetAllLimit, 0).Return(nil, nil)

	srv.GetAll()
}

func TestDeviceService_AddNew(t *testing.T) {
	repo := interfaces.MockDeviceRepository{}
	srv := NewDeviceService(&repo)

	deviceCreate := model.Device{
		Name:  "nm1",
		Brand: "br1",
	}
	deviceCreated := model.Device{
		Id:      123,
		Name:    deviceCreate.Name,
		Brand:   deviceCreate.Brand,
		Created: time.Now(),
	}
	repo.On("Add", &deviceCreate).Return(&deviceCreated, nil)

	device, err := srv.AddNew(&deviceCreate)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if device != &deviceCreated {
		t.Error("created device is not as expected ", deviceCreated)
	}
}

func TestDeviceService_AddNew_error(t *testing.T) {
	repo := interfaces.MockDeviceRepository{}
	srv := NewDeviceService(&repo)

	deviceCreate := model.Device{
		Name:  "nm1",
		Brand: "br1",
	}
	addError := errors.New("unable to create")
	repo.On("Add", &deviceCreate).Return(nil, addError)

	device, err := srv.AddNew(&deviceCreate)
	if !errors.Is(err, addError) {
		t.Error("error should be returned", err)
	}
	if device != nil {
		t.Error("error generated but device also returned")
	}
}
