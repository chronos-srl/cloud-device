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
	info *device.Info
}

func NewEmt130() device.Device {
	return Emt130{
		info: &device.Info{
			Name:    "Emt130",
			Version: "v1.0.0",
		},
	}
}

func (e Emt130) GetId() string {
	return "emt130:v1"
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
		if err := mapping.UnmarshalUint16(response.Values, regs); err != nil {
			return nil, err
		}

		return response, nil

	default:
		return "", errors.New("not implemented")
	}
}

func (e Emt130) GetWriteRequestBytes(ctx context.Context, body []byte) (command.DeviceWriteRequest, error) {
	//TODO implement me
	panic("implement me")
}
