package repository

import (
	"database/sql"
	"devices/interfaces"
	"devices/internal/model"
	"errors"
)

type DeviceRepositoryImpl struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) interfaces.DeviceRepository {
	return &DeviceRepositoryImpl{db: db}
}

func (d *DeviceRepositoryImpl) List(limit, offset int) ([]*model.Device, error) {
	devices := make([]*model.Device, 0)

	rows, err := d.db.Query("SELECT id, name, brand, created FROM device LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device model.Device
		if err := rows.Scan(&device.Id, &device.Name, &device.Brand, &device.Created); err != nil {
			return nil, err
		}
		devices = append(devices, &device)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

func (d *DeviceRepositoryImpl) GetById(id int64) (*model.Device, error) {
	var device model.Device
	row := d.db.QueryRow("SELECT * FROM device WHERE id = ?", id)
	if err := row.Scan(&device.Id, &device.Name, &device.Brand, &device.Created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &device, nil
}

func (d *DeviceRepositoryImpl) Add(device *model.Device) (*model.Device, error) {
	result, err := d.db.Exec("INSERT INTO device (name, brand) VALUES (?, ?)", device.Name, device.Brand)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return d.GetById(id)
}

func (d *DeviceRepositoryImpl) Update(device *model.Device) (*model.Device, error) {
	_, err := d.db.Exec("UPDATE device SET name = ?, brand = ? WHERE id = ?", device.Name, device.Brand, device.Id)
	if err != nil {
		return nil, err
	}

	return d.GetById(device.Id)
}

func (d *DeviceRepositoryImpl) Delete(id int64) error {
	_, err := d.db.Exec("DELETE FROM device WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DeviceRepositoryImpl) SearchByBrand(brand string) ([]*model.Device, error) {
	devices := make([]*model.Device, 0)

	rows, err := d.db.Query("SELECT id, name, brand, created FROM device WHERE brand LIKE ?", brand+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device model.Device
		if err := rows.Scan(&device.Id, &device.Name, &device.Brand, &device.Created); err != nil {
			return nil, err
		}
		devices = append(devices, &device)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}
