.PHONY: build
build:
	cd cmd/nfsauto; go build -ldflags "-s -X main.Version=0.0.0" -o ../../bin/nfsauto;
	cd cmd/volumeauto; go build -o ../../bin/volumeauto;

.PHONY: exec
exec: build
	cd cmd/nfsauto; sudo ./bin/nfsauto
