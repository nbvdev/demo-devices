package interfaces

import "devices/internal/model"

type DeviceRepository interface {
	List(limit, offset int) ([]*model.Device, error)
	SearchByBrand(brand string) ([]*model.Device, error)
	GetById(id int64) (*model.Device, error)
	Add(model *model.Device) (*model.Device, error)
	Update(model *model.Device) (*model.Device, error)
	Delete(id int64) error
}
