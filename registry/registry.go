package registry

import (
	"github.com/chronos-srl/cloud-device/device"
)

// Registry save all known devices
type Registry interface {
	// Add new device to registry. Device must be added to be used
	Add(device device.Device) error
	// Exists check if device id is already registered
	Exists(device device.Device) bool
	// ExistsId check if device id is already registered
	ExistsId(id string) bool
	// Get device by device id, throw error if not found
	Get(id string) (device.Device, error)
	// GetAllInfo return all device information
	GetAllInfo() ([]device.Info, error)
}
