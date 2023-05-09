module github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway

go 1.20

replace github.com/plaoludastruja/JBSPLS/Skitnica/backend/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/plaoludastruja/JBSPLS/Skitnica/backend/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/rs/cors v1.9.0
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
