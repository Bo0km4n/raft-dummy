pb:
	protoc --go_out=plugins=grpc:. ./proto/*.proto
test:
	GOCACHE=off go test ./... -v -cwd="$$PWD"
