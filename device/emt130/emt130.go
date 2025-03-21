package emt130

import (
	"context"
	"errors"
	"github.com/chronos-srl/cloud-device/device"
	"github.com/chronos-srl/cloud-protocol/command"
	"github.com/chronos-srl/cloud-protocol/mapping"
	"slices"
)

var (
	_ device.Device = (*Emt130)(nil)
)

type Emt130 struct {
	info            *device.Info
	metricsRequests []command.ReadRegistryRequest
}

func NewEmt130() device.Device {
	metricsRequests := make([]command.ReadRegistryRequest, 0)
	mr, _ := command.NewRegistryReadRequest(&Metrics{})
	metricsRequests = append(metricsRequests, mr)
	mr2, _ := command.NewRegistryReadRequest(&Metrics2{})
	metricsRequests = append(metricsRequests, mr2)

	return Emt130{
		info: &device.Info{
			Model:           "emt-130",
			FirmwareVersion: "1.0.0",
		},
		metricsRequests: metricsRequests,
	}
}

func (e Emt130) GetMetricsRequests(_ context.Context) ([]command.ReadRegistryRequest, error) {
	return e.metricsRequests, nil
}

func (e Emt130) GetModel() string {
	return e.info.Model
}

func (e Emt130) GetInfo() *device.Info {
	return e.info
}

func (e Emt130) ParseMetricsRequest(_ context.Context, index int, response command.ReadDeviceRegistryResponse) (mapping.ValueMap, error) {
	bytes := mapping.ToBytes(response.Values)
	switch index {
	case 0:
		return mapping.AsValueMapTyped(bytes, Metrics{})
	case 1:
		return mapping.AsValueMapTyped(bytes, Metrics2{})
	}

	return nil, errors.New("not implemented")
}

func (e Emt130) GetRegistries(_ context.Context) (mapping.Registries, error) {
	regs1, err := mapping.AsRegistries(Metrics{})
	if err != nil {
		return nil, err
	}

	regs2, err := mapping.AsRegistries(Metrics2{})
	if err != nil {
		return nil, err
	}

	return slices.Concat(regs1, regs2), nil
}
