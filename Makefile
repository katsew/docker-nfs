.PHONY: build
build:
	cd cmd/nfsauto; go build -o tmp/nfsauto;

.PHONY: exec
exec: build
	cd cmd/nfsauto; sudo ./tmp/nfsauto
