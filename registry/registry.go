package registry

import (
	"cloud-device/device"
)

type Registry interface {
	Add(device device.Device) error
	Exists(device device.Device) bool
	ExistsId(id string) bool
	Get(deviceType string) (device.Device, error)
	GetAllInfo() ([]device.Info, error)
}
