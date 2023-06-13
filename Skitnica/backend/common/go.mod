module github.com/plaoludastruja/JBSPLS/Skitnica/backend/common

go 1.20

replace github.com/plaoludastruja/JBSPLS/Skitnica/backend/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	go.mongodb.org/mongo-driver v1.11.7
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/nats-io/nats-server/v2 v2.9.17 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.8.0 // indirect
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/nats-io/nats.go v1.26.0
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
