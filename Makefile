nfsautoVersion = 1.0.0
volumeautoVersion = 1.0.0

.PHONY: nfsauto
nfsauto:
	cd cmd/nfsauto; go build -ldflags "-s -X main.Version=$(nfsautoVersion)" -o ../../bin/nfsauto;

.PHONY: volumeauto
volumeauto:
	cd cmd/volumeauto; go build -ldflags "-s -X main.Version=$(volumeautoVersion)" -o ../../bin/volumeauto;

.PHONY: build
build: nfsauto volumeauto
