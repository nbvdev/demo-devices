package interfaces

import "devices/internal/model"

type DeviceService interface {
	GetAll() ([]*model.Device, error)
	AddNew(device *model.Device) (*model.Device, error)
	Get(deviceId int64) (*model.Device, error)
	Update(device *model.Device) (*model.Device, error)
	UpdatePartial(device *model.Device) (*model.Device, error)
	Delete(deviceId int64) error
	SearchByBrand(brand string) ([]*model.Device, error)
}
