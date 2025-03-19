package emt130

type Metrics struct {
	Frequency float32 `json:"frequency" addr:"1"`
	Ampere    float32 `json:"ampere"`
}

type Metrics2 struct {
	Vln float32 `json:"vln" addr:"430"`
}
