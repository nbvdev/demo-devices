package service

import (
	"devices/interfaces"
	"devices/internal/model"
	"errors"
)

const GetAllLimit = 1000

type DeviceServiceImpl struct {
	repository interfaces.DeviceRepository
}

func NewDeviceService(repository interfaces.DeviceRepository) interfaces.DeviceService {
	return &DeviceServiceImpl{repository}
}

func (s *DeviceServiceImpl) GetAll() ([]*model.Device, error) {
	return s.repository.List(GetAllLimit, 0)
}

func (s *DeviceServiceImpl) AddNew(device *model.Device) (*model.Device, error) {
	return s.repository.Add(device)
}

func (s *DeviceServiceImpl) Get(deviceId int64) (*model.Device, error) {
	return s.repository.GetById(deviceId)
}

func (s *DeviceServiceImpl) Update(device *model.Device) (*model.Device, error) {
	if device.IsSuitableForUpdate() == false {
		return nil, errors.New("not suitable for update")
	}

	return s.UpdatePartial(device)
}

func (s *DeviceServiceImpl) UpdatePartial(device *model.Device) (*model.Device, error) {
	currentDevice, err := s.repository.GetById(device.Id)
	if err != nil {
		return nil, err
	}
	if currentDevice == nil {
		return nil, nil
	}
	currentDevice.Patch(device)

	return s.repository.Update(currentDevice)
}

func (s *DeviceServiceImpl) Delete(deviceId int64) error {
	return s.repository.Delete(deviceId)
}

func (s *DeviceServiceImpl) SearchByBrand(brand string) ([]*model.Device, error) {
	return s.repository.SearchByBrand(brand)
}
