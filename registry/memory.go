package registry

import (
	"errors"
	"github.com/chronos-srl/cloud-device/device"
	"github.com/chronos-srl/cloud-device/device/emt130"
	"github.com/chronos-srl/cloud-device/device/oven"
)

var (
	_ Registry = (*memoryRegistry)(nil)
)

// NewMemoryRegistry create in-memory device registry
func NewMemoryRegistry() Registry {
	return &memoryRegistry{
		devices: make(map[string]device.Device),
	}
}

type memoryRegistry struct {
	devices map[string]device.Device
}

func (m *memoryRegistry) Load() {
	_ = m.Add(emt130.NewEmt130())
	_ = m.Add(oven.NewOven())
}

func (m *memoryRegistry) Add(device device.Device) error {
	if device == nil {
		return errors.New("cannot add nil device")
	}

	if m.ExistsId(device.GetModel()) {
		return NewDeviceError(device.GetModel(), "device with this id already exists")
	}

	m.devices[device.GetModel()] = device
	return nil
}

func (m *memoryRegistry) Exists(device device.Device) bool {
	return m.ExistsId(device.GetModel())
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

func (m *memoryRegistry) GetAllInfo() ([]device.Info, error) {
	info := make([]device.Info, 0)
	for _, s := range m.devices {
		info = append(info, *s.GetInfo())
	}

	return info, nil
}
