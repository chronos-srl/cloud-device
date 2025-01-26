package emt130

import (
	"context"
	"errors"
	"github.com/chronos-srl/cloud-device/device"
	"github.com/chronos-srl/cloud-protocol/command"
	"github.com/chronos-srl/cloud-protocol/mapping"
)

var (
	_ device.Device = (*Emt130)(nil)
)

type Emt130 struct {
	info            *device.Info
	metricsRequests []command.RegistryReadRequest
}

func NewEmt130() device.Device {
	metricsRequests := make([]command.RegistryReadRequest, 0)
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

func (e Emt130) GetMetricsRequests(ctx context.Context) ([]command.RegistryReadRequest, error) {
	return e.metricsRequests, nil
}

func (e Emt130) GetModel() string {
	return e.info.Model
}

func (e Emt130) GetInfo() *device.Info {
	return e.info
}

func (e Emt130) GetReadRequest(rt command.RequestType) (command.DeviceReadRequest, error) {
	return command.DeviceReadRequest{}, errors.New("not implemented")
}

func (e Emt130) ParseReadRequest(ctx context.Context, rt command.RequestType, response command.ReadResponse) (interface{}, error) {
	switch rt {
	case command.ReadMetricsType:
		var regs = new(Metrics)
		if err := mapping.Unmarshal(response.Values, regs); err != nil {
			return nil, err
		}

		return response, nil

	default:
		return "", errors.New("not implemented")
	}
}

func (e Emt130) ParseMetricsRequest(ctx context.Context, index int, response command.ReadResponse) (mapping.ValueMap, error) {
	switch index {
	case 0:
		return mapping.AsValueMapTyped(response.Values, Metrics{})
	case 1:
		return mapping.AsValueMapTyped(response.Values, Metrics2{})
	}

	return nil, errors.New("not implemented")
}

func (e Emt130) GetWriteRequestBytes(ctx context.Context, body []byte) (command.DeviceWriteRequest, error) {
	panic("implement me")
}

func (e Emt130) GetRegistries(ctx context.Context) (mapping.Registries, error) {
	return mapping.AsRegistries(Metrics{})
}
