package device

import (
	"context"
	"github.com/chronos-srl/cloud-protocol/command"
	"github.com/chronos-srl/cloud-protocol/mapping"
)

// Device a physical device
type Device interface {
	// GetModel unique device identifier
	GetModel() string
	// GetInfo device information name, firmware ...
	GetInfo() *Info
	// GetMetricsRequests return all metrics request for device
	GetMetricsRequests(ctx context.Context) ([]command.ReadRegistryRequest, error)
	// ParseMetricsRequest parse device metrics request into a `ValueMap` as map[string]any
	ParseMetricsRequest(ctx context.Context, index int, response command.ReadDeviceRegistryResponse) (mapping.ValueMap, error)
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
