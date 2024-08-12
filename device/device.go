package device

import (
	"context"
	"github.com/chronos-srl/cloud-protocol/command"
)

// Device a physical device
type Device interface {
	// GetId unique device identifier
	GetId() string
	// GetInfo device information name, firmware ...
	GetInfo() *Info
	// GetReadRequest return the read request that indicate the registry address and quantity to read from device
	GetReadRequest(rt command.RequestType) (command.DeviceReadRequest, error)
	// ParseReadRequest parse device response into a device interface.
	// In this function we can do some modification and return a frontend ready struct
	ParseReadRequest(ctx context.Context, rt command.RequestType, response command.ReadResponse) (interface{}, error)
	// GetWriteRequestBytes encode the incoming frontend request into a modbus registry payload
	// This function convert the frontend struct into a modbus bytes ready values
	GetWriteRequestBytes(ctx context.Context, body []byte) (command.DeviceWriteRequest, error)
}

type BaseDevice struct {
	Type string
	Info *Info
}

// Info describe Device information
type Info struct {
	Name             string `json:"name"`
	Model            string `json:"model"`
	SerialNumber     string `json:"serialNumber"`
	FirmwareRevision string `json:"firmwareRevision"`
}
