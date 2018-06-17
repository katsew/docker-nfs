.PHONY: init
init:
	go install github.com/katsew/docker-nfs/cmd/nfsauto
	go install github.com/katsew/docker-nfs/cmd/volumeauto
	sudo nfsauto
	cd public; volumeauto -name docker-nfs_example_public

# default volume option: addr=host.docker.internal,rw
.PHONY: up
up:
	cd public; volumeauto -name docker-nfs_example_public -f
	docker-compose up -d
	docker-compose ps
	open http://localhost:9090

.PHONY: halt
halt:
	docker-compose kill && docker-compose rm -f
	docker-compose ps
	docker volume rm docker-nfs_example_public
	docker volume ls

.PHONY: reload
reload: halt up

.PHONY: volume
volume:
	cd public; volumeauto -name docker-nfs_example_public -f
	docker inspect docker-nfs_example_public
