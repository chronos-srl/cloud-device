package registry

import (
	"errors"
	"github.com/chronos-srl/cloud-device/device"
)

var (
	_ Registry = (*memoryRegistry)(nil)
)

type memoryRegistry struct {
	devices map[string]device.Device
}

// NewMemoryRegistry create in-memory device registry
func NewMemoryRegistry() Registry {
	return &memoryRegistry{
		devices: make(map[string]device.Device),
	}
}

func (m *memoryRegistry) Add(device device.Device) error {
	if device == nil {
		return errors.New("cannot add nil device")
	}

	if m.ExistsId(device.GetId()) {
		return NewDeviceError(device.GetId(), "device with this id already exists")
	}

	m.devices[device.GetId()] = device
	return nil
}

func (m *memoryRegistry) Exists(device device.Device) bool {
	return m.ExistsId(device.GetId())
}

func (m *memoryRegistry) ExistsId(id string) bool {
	_, ok := m.devices[id]
	return ok
}

func (m *memoryRegistry) Get(id string) (device.Device, error) {
	d, ok := m.devices[id]
	if !ok {
		return nil, NewDeviceError(id, "device not found")
	}

	return d, nil
}

func (m *memoryRegistry) GetAllInfo() ([]device.BaseDevice, error) {
	devices := make([]device.BaseDevice, 0)
	for _, s := range m.devices {
		devices = append(devices, device.BaseDevice{
			Id:   s.GetId(),
			Info: s.GetInfo(),
		})
	}

	return devices, nil
}
