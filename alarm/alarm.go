package alarm

type DeviceAlarm struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
}
