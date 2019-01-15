pb:
	protoc --go_out=plugins=grpc:. ./proto/*.proto
test:
	go test ./... -v -cwd="$$PWD"
