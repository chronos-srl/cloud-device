package oven

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/chronos-srl/cloud-device/device"
	"github.com/chronos-srl/cloud-protocol/command"
)

var (
	_ device.Device = (*Oven)(nil)
)

type Oven struct {
	device.BaseDevice
	paramsRequest command.DeviceReadRequest
	alarmsRequest command.DeviceReadRequest
}

var alarmsString = []string{
	"TFT Comunication Error",
	"Board Temp Val out of Range",
	"CN1 Temp Val out of Range",
	"CN3 Temp Val out of Range",
	"CN5 Temp Val out of Range",
}

type ovenParamRegistry struct {
	TemperaturaSchedaCarichi   uint16 `json:"temperaturaSchedaCarichi"`
	TemperaturaCN1             uint16 `json:"temperaturaCN1" swap:"ba"`
	TemperaturaCN3             uint16 `json:"temperaturaCN3"`
	TemperaturaCN5             uint16 `json:"temperaturaCN5"`
	IngressiDigitaliCN12       uint16 `json:"ingressiDigitaliCN12"`
	ErrorFlags                 uint16 `json:"errorFlags"`
	Coils                      uint16 `json:"coils"`
	TemperaturaImpostataCielo  uint16 `json:"temperaturaImpostataCielo"`
	TemperaturaImpostataPlatea uint16 `json:"temperaturaImpostataPlatea"`
	PotenzaImpostataCielo      uint16 `json:"potenzaImpostataCielo"`
	PotenzaImpostataPlatea     uint16 `json:"potenzaImpostataPlatea"`
	OperativeFlags             uint16 `json:"operativeFlags"`
}

type ovenParamResponse struct {
	TemperaturaSchedaCarichi   uint16             `json:"temperaturaSchedaCarichi"`
	TemperaturaCN1             uint16             `json:"temperaturaCN1"`
	TemperaturaCN3             uint16             `json:"temperaturaCN3"`
	TemperaturaCN5             uint16             `json:"temperaturaCN5"`
	IngressiDigitaliCN12       uint16             `json:"ingressiDigitaliCN12"`
	ErrorFlags                 uint16             `json:"errorFlags"`
	Coils                      uint16             `json:"coils"`
	TemperaturaImpostataCielo  uint16             `json:"temperaturaImpostataCielo"`
	TemperaturaImpostataPlatea uint16             `json:"temperaturaImpostataPlatea"`
	PotenzaImpostataCielo      uint16             `json:"potenzaImpostataCielo"`
	PotenzaImpostataPlatea     uint16             `json:"potenzaImpostataPlatea"`
	OperativeFlags             ovenOperativeFlags `json:"operativeFlags"`
	Errors                     []string           `json:"errors"`
}

type ovenOperativeFlags struct {
	BoccaOn                bool `json:"boccaOn"`
	Booster                bool `json:"booster"`
	CappaOn                bool `json:"cappaOn"`
	Celsius                bool `json:"celsius"`
	DualTemp               bool `json:"dualTemp"`
	DualTemp2              bool `json:"dualTemp2"`
	EcoMode                bool `json:"ecoMode"`
	Formato24h             bool `json:"formato24h"`
	PlateaOff              bool `json:"plateaOff"`
	PowerBoost             bool `json:"powerBoost"`
	PowerOn                bool `json:"powerOn"`
	PowerOnTimer           bool `json:"powerOnTimer"`
	ResistenzaBocca        bool `json:"resistenzaBocca"`
	RichiestaRallentamento bool `json:"richiestaRallentamento"`
	Running                bool `json:"running"`
	TermocoppiaJ           bool `json:"termocoppiaJ"`
}

type ovenWriteRegistry struct {
	TemperaturaImpostataCielo  uint16 `json:"temperaturaImpostataCielo"`
	TemperaturaImpostataPlatea uint16 `json:"temperaturaImpostataPlatea"`
	PotenzaImpostataCielo      uint16 `json:"potenzaImpostataCielo"`
	PotenzaImpostataPlatea     uint16 `json:"potenzaImpostataPlatea"`
}

type AlarmRegistry struct {
	ErrorFlags uint16 `addr:"5"`
}

type ovenMetricsRegistry struct {
	TemperaturaSchedaCarichi uint16 `json:"temperaturaSchedaCarichi"`
	TemperaturaCN1           uint16 `json:"temperaturaCN1" swap:"ba"`
}

func NewOven() *Oven {
	d := device.BaseDevice{
		Type: "oven-01",
		Info: &device.Info{
			Name:             "Forno",
			Model:            "Model-01",
			SerialNumber:     "Serial-01",
			FirmwareRevision: "1.0.0",
		},
	}

	paramsRequest, _ := command.NewDeviceReadRequest(&ovenParamRegistry{})
	alarmsRequest, _ := command.NewDeviceReadRequest(&AlarmRegistry{})

	return &Oven{
		BaseDevice:    d,
		paramsRequest: paramsRequest,
		alarmsRequest: alarmsRequest,
	}
}

func (o *Oven) GetId() string {
	return o.BaseDevice.Type
}

func (o *Oven) GetInfo() *device.Info {
	return o.Info
}

func (o *Oven) GetReadRequest(rt command.RequestType) (command.DeviceReadRequest, error) {
	switch rt {
	case command.ReadParamsType:
		return o.paramsRequest, nil

	case command.ReadAlarmsType:
		return o.alarmsRequest, nil

	default:
		return command.DeviceReadRequest{}, errors.New("invalid request type")
	}
}

func (o *Oven) ParseReadRequest(ctx context.Context, rt command.RequestType, response command.ReadResponse) (interface{}, error) {
	switch rt {
	case command.ReadParamsType:
		v := new(ovenParamRegistry)
		if err := command.ParseReadResponse(response, v); err != nil {
			return nil, err
		}

		errs := make([]string, 0)
		for j := 0; j < 16; j++ {
			v1 := v.ErrorFlags >> j & 01
			if v1 == 1 {
				errs = append(errs, alarmsString[j])
			}
		}

		fv := ovenParamResponse{
			TemperaturaSchedaCarichi:   v.TemperaturaSchedaCarichi,
			TemperaturaCN1:             v.TemperaturaCN1,
			TemperaturaCN3:             v.TemperaturaCN3,
			TemperaturaCN5:             v.TemperaturaCN5,
			IngressiDigitaliCN12:       v.IngressiDigitaliCN12,
			ErrorFlags:                 v.ErrorFlags,
			Coils:                      v.Coils,
			TemperaturaImpostataCielo:  v.TemperaturaImpostataCielo,
			TemperaturaImpostataPlatea: v.TemperaturaImpostataPlatea,
			PotenzaImpostataCielo:      v.PotenzaImpostataCielo,
			PotenzaImpostataPlatea:     v.PotenzaImpostataPlatea,
			Errors:                     errs,
			OperativeFlags: ovenOperativeFlags{
				PowerBoost:             v.OperativeFlags&0x1 == 1,
				EcoMode:                v.OperativeFlags>>1&0x1 == 1,
				ResistenzaBocca:        v.OperativeFlags>>2&0x1 == 1,
				RichiestaRallentamento: v.OperativeFlags>>3&0x1 == 1,
				Booster:                v.OperativeFlags>>4&0x1 == 1,
				PlateaOff:              v.OperativeFlags>>5&0x1 == 1,
				DualTemp:               v.OperativeFlags>>6&0x1 == 1,
				PowerOn:                v.OperativeFlags>>7&0x1 == 1,
				DualTemp2:              v.OperativeFlags>>8&0x1 == 1,
				PowerOnTimer:           v.OperativeFlags>>9&0x1 == 1,
				Formato24h:             v.OperativeFlags>>10&0x1 == 1,
				Celsius:                v.OperativeFlags>>11&0x1 == 1,
				TermocoppiaJ:           v.OperativeFlags>>12&0x1 == 1,
				CappaOn:                v.OperativeFlags>>13&0x1 == 1,
				BoccaOn:                v.OperativeFlags>>14&0x1 == 1,
				Running:                v.OperativeFlags>>15&0x1 == 1,
			},
		}

		return fv, nil

	case command.ReadAlarmsType:
		v := new(AlarmRegistry)
		if err := command.ParseReadResponse(response, v); err != nil {
			return nil, err
		}

		return v, nil

	case command.ReadMetricsType:
		v := new(ovenMetricsRegistry)
		if err := command.ParseReadResponse(response, v); err != nil {
			return nil, err
		}

		var f interface{}
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &f); err != nil {
			return nil, err
		}

		return f.(map[string]interface{}), err

	default:
		return nil, errors.New("invalid request type")
	}
}

func (o *Oven) GetWriteRequestBytes(ctx context.Context, body []byte) (command.DeviceWriteRequest, error) {
	v := new(ovenWriteRegistry)
	d := command.DeviceWriteRequest{}

	if err := json.Unmarshal(body, v); err != nil {
		return d, errors.New("cannot decode payload")
	}

	return command.NewDeviceWriteRequest(v)
}
