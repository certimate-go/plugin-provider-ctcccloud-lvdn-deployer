.PHONY: build test

build:
	GOWORK=off go build -mod=readonly -o ctcccloud-lvdn.exe .

test:
	GOWORK=off go test -mod=readonly ./...
