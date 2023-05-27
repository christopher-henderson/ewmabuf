module github.com/whatever

go 1.20

require (
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

replace google.golang.org/protobuf => github.com/chris-henderson-alation/protobuf-go v0.0.0-20230524180019-dbb5c83f837f

//replace google.golang.org/grpc => github.com/chris-henderson-alation/grpc-go v0.0.0-20230524210947-9c189242fdb0
