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
	info           *device.Info
	metricsRequest command.DeviceReadRequest
}

func NewEmt130() device.Device {
	metricsRequest, _ := command.NewDeviceReadRequest(&Metrics{})
	return Emt130{
		info: &device.Info{
			Model:           "emt-130",
			FirmwareVersion: "1.0.0",
		},
		metricsRequest: metricsRequest,
	}
}

func (e Emt130) GetModel() string {
	return e.info.Model
}

func (e Emt130) GetInfo() *device.Info {
	return e.info
}

func (e Emt130) GetReadRequest(rt command.RequestType) (command.DeviceReadRequest, error) {
	switch rt {
	case command.ReadMetricsType:
		return e.metricsRequest, nil
	default:
		return command.DeviceReadRequest{}, errors.New("not implemented")
	}
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

func (e Emt130) ParseMetricsRequest(ctx context.Context, response command.ReadResponse) (mapping.ValueMap, error) {
	return mapping.AsValueMapTyped(response.Values, Metrics{})
}

func (e Emt130) GetWriteRequestBytes(ctx context.Context, body []byte) (command.DeviceWriteRequest, error) {
	panic("implement me")
}

func (e Emt130) GetRegistries(ctx context.Context) (mapping.Registries, error) {
	return mapping.AsRegistries(Metrics{})
}
