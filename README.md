# docker-nfs

Auto configuration cli for insert docker nfs config.
Command will do the following actions:

1. Insert config (or create new) `/etc/exports`
2. Insert config (or create new) `/etc/nfs.conf`
3. Restart `nfsd` if at least one of the file appears above updates

:zap: Use this command with `sudo` or your config file will accidentally update!

## Installation

```bash
$ go install github.com/katsew/docker-nfs/cmd/nfsauto
```

## Execute

```bash
$ sudo nfsauto -addr 192.168.33.10 -volume project_local
```

e.g.

```
$ cd /path/to/your/docker/project
$ sudo nfsauto 192.168.33.10
$ cat /etc/exports

# BEGIN - docker-nfs [uid]:[gid]
"/path/to/your/docker/project" 192.168.33.10 -alldirs -mapall=[uid]:[gid]
# END - docker-nfs [uid]:[gid]

$ cat /etc/nfs.conf

## BEGIN - docker-nfs
nfs.server.mount.require_resv_port = 0
## END - docker-nfs

nfsd will be restarted

$ docker volume inspect project_local

[
    {
        "CreatedAt": "",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/project_local/_data",
        "Name": "project_local",
        "Options": {
            "device": ":/path/to/your/docker/project",
            "o": "addr=host.docker.internal,rw,vers=3,tcp,fsc,actimeo=2",
            "type": "nfs"
        },
        "Scope": "local"
    }
]

```

## License

MIT
