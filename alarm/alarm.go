package alarm

type HasAlarm interface {
	GetAlarms() ([]Alarm, error)
}

type Alarm struct {
	Message string `json:"message"`
	Code    uint16 `json:"code"`
}
