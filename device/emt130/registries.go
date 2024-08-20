package emt130

type Metrics struct {
	Vl1n      float32 `json:"vl1n" addr:"40359"`
	Vl2n      float32 `json:"vl2n"`
	Vl3n      float32 `json:"vl3n"`
	VstarAvg  float32 `json:"vstarAvg"`
	Vl1l2     float32 `json:"vl1l2"`
	Vl2l3     float32 `json:"vl2l3"`
	Vl3l1     float32 `json:"vl3l1"`
	VlineAvg  float32 `json:"vlineAvg"`
	Il1       float32 `json:"il1"`
	Il2       float32 `json:"il2"`
	Il3       float32 `json:"il3"`
	In        float32 `json:"in"`
	Iavg      float32 `json:"iavg"`
	P1        float32 `json:"p1"`
	P2        float32 `json:"p2"`
	P3        float32 `json:"p3"`
	Psum      float32 `json:"psum"`
	Q1        float32 `json:"q1"`
	Q2        float32 `json:"q2"`
	Q3        float32 `json:"q3"`
	Qsum      float32 `json:"qsum"`
	S1        float32 `json:"s1"`
	S2        float32 `json:"s2"`
	S3        float32 `json:"s3"`
	Ssum      float32 `json:"ssum"`
	Pf1       float32 `json:"pf1"`
	Pf2       float32 `json:"pf2"`
	Pf3       float32 `json:"pf3"`
	Pf3ph     float32 `json:"pf3h"`
	Cf1       float32 `json:"cf1"`
	Cf2       float32 `json:"cf2"`
	Cf3       float32 `json:"cf3"`
	Cfn       float32 `json:"cfn"`
	Frequency float32 `json:"frequency"`
	Vl1npeak  float32 `json:"vl1npeak"`
	Vl2npeak  float32 `json:"vl2npeak"`
	Vl3npeak  float32 `json:"vl3npeak"`
	Vl1l2peak float32 `json:"vl1l2peak"`
	Vl2l3peak float32 `json:"vl2l3peak"`
	Vl3l1peak float32 `json:"vl3l1peak"`
	Il1peak   float32 `json:"il1peak"`
	Il2peak   float32 `json:"il2peak"`
	Il3peak   float32 `json:"il3peak"`
	Inpeak    float32 `json:"inpeak"`
}
