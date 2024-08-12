package registry

type DeviceError struct {
	Message string
	Id      string
}

func (e DeviceError) Error() string {
	return e.Message
}

func NewDeviceError(id string, msg string) error {
	return DeviceError{
		Message: msg,
		Id:      id,
	}
}
