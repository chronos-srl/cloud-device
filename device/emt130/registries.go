package emt130

type VLMetrics struct {
	VL1n float32 `json:"vl1n" addr:"40833"`
	VL2n float32 `json:"vl2n"`
	Vl3n float32 `json:"vl3n"`
}

type IMetrics struct {
	IL1  float32 `json:"il1" addr:"40849"`
	IL2  float32 `json:"il2"`
	IL3  float32 `json:"il3"`
	IN   float32 `json:"in"`
	Iavg float32 `json:"iavg"`
	P1   float32 `json:"p1"`
	P2   float32 `json:"p2"`
	P3   float32 `json:"p3"`
	PSum float32 `json:"pSum"`
	Q1   float32 `json:"q1"`
	Q2   float32 `json:"q2"`
	Q3   float32 `json:"q3"`
	QSum float32 `json:"qSum"`
	Pf1  float32 `json:"pf1"`
	Pf2  float32 `json:"pf2"`
	Pf3  float32 `json:"pf3"`
}
