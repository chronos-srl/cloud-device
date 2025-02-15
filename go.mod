module github.com/chronos-srl/cloud-device

go 1.23

replace github.com/chronos-srl/cloud-protocol => ../cloud-protocol

require github.com/chronos-srl/cloud-protocol v1.1.1-0.20250128092107-c03ffed310c1

require (
	github.com/fxamacker/cbor v1.5.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
