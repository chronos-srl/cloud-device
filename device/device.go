package device

import (
	"context"
	"github.com/chronos-srl/cloud-protocol/mapping"
	"github.com/chronos-srl/cloud-protocol/pb"
)

// Device a physical device
type Device interface {
	// GetModel unique device identifier
	GetModel() string
	// GetInfo device information name, firmware ...
	GetInfo() *Info
	// GetReadRequest return the read request that indicate the registry address and quantity to read from device
	GetReadRequest() (pb.ReadRegistryRequest, error)
	// GetMetricsRequests return all metrics request for device
	GetMetricsRequests(ctx context.Context) ([]pb.ReadRegistryRequest, error)
	// ParseReadRequest parse device response into a device interface.
	// In this function we can do some modification and return a frontend ready struct
	ParseReadRequest(ctx context.Context, response *pb.ReadRegistryResponse) (interface{}, error)
	// ParseMetricsRequest parse device metrics request into a `ValueMap` as map[string]any
	ParseMetricsRequest(ctx context.Context, index int, response *pb.DeviceReadRegistryResponse) (mapping.ValueMap, error)
	// GetWriteRequestBytes encode the incoming frontend request into a modbus registry payload
	// This function convert the frontend struct into a modbus bytes ready values
	GetWriteRequestBytes(ctx context.Context, body []byte) (*pb.WriteRegistryRequest, error)
	// GetRegistries obtain all metrics request registries
	GetRegistries(ctx context.Context) (mapping.Registries, error)
}

type BaseDevice struct {
	Info *Info
}

// Info describe Device information
type Info struct {
	Model           string `json:"model"`
	FirmwareVersion string `json:"firmwareVersion"`
}
