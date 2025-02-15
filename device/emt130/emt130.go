package emt130

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/chronos-srl/cloud-device/device"
	"github.com/chronos-srl/cloud-protocol/command"
	"github.com/chronos-srl/cloud-protocol/mapping"
	"github.com/chronos-srl/cloud-protocol/pb"
)

var (
	_ device.Device = (*Emt130)(nil)
)

type Emt130 struct {
	info            *device.Info
	metricsRequests []pb.ReadRegistryRequest
}

func NewEmt130() device.Device {
	metricsRequests := make([]pb.ReadRegistryRequest, 0)
	mr, _ := command.MappingStructToReadRegistry(&VLMetrics{})
	metricsRequests = append(metricsRequests, mr)
	mr2, _ := command.MappingStructToReadRegistry(&IMetrics{})
	metricsRequests = append(metricsRequests, mr2)

	return Emt130{
		info: &device.Info{
			Model:           "emt-130",
			FirmwareVersion: "1.0.0",
		},
		metricsRequests: metricsRequests,
	}
}

func (e Emt130) GetReadRequest() (pb.ReadRegistryRequest, error) {
	return pb.ReadRegistryRequest{}, errors.New("not implemented")
}

func (e Emt130) GetMetricsRequests(_ context.Context) ([]pb.ReadRegistryRequest, error) {
	return e.metricsRequests, nil
}

func (e Emt130) ParseReadRequest(_ context.Context, _ *pb.ReadRegistryResponse) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (e Emt130) ParseMetricsRequest(_ context.Context, index int, response *pb.DeviceReadRegistryResponse) (mapping.ValueMap, error) {
	// TODO cambiare metodo AsValueMapTyped
	rBytes := response.Response.Values
	v := make([]uint16, len(rBytes)/2)
	for i := 0; i < len(v); i += 2 {
		v[i] = binary.BigEndian.Uint16(rBytes[i : i+2])
	}

	switch index {
	case 0:
		return mapping.AsValueMapTyped(v, VLMetrics{})
	case 1:
		return mapping.AsValueMapTyped(v, IMetrics{})
	}

	return nil, errors.New("not implemented")
}

func (e Emt130) GetModel() string {
	return e.info.Model
}

func (e Emt130) GetInfo() *device.Info {
	return e.info
}

func (e Emt130) GetWriteRequestBytes(_ context.Context, _ []byte) (*pb.WriteRegistryRequest, error) {
	panic("implement me")
}

func (e Emt130) GetRegistries(_ context.Context) (mapping.Registries, error) {
	return mapping.AsRegistries(VLMetrics{})
}
