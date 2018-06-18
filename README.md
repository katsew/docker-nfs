# docker-nfs

[![CircleCI](https://circleci.com/gh/katsew/docker-nfs/tree/master.svg?style=svg)](https://circleci.com/gh/katsew/docker-nfs/tree/master)

configuration automation tools for docker nfs.

## nfsauto

Auto configuration cli for native nfs configuration.
Command will do the following actions:

1. Insert config (or create new) `/etc/exports`
2. Insert config (or create new) `/etc/nfs.conf`
3. Restart `nfsd` if at least one of the file appears above updates

:zap: Use this command with `sudo` or your config file will accidentally update!

### Installation

```bash
$ go get github.com/katsew/docker-nfs/cmd/nfsauto
```

### Execute

```bash
$ sudo nfsauto -addr 192.168.33.10
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

```

## volumeauto

Create docker volume with type nfs for current directory

### Installation

```bash
$ go get github.com/katsew/docker-nfs/cmd/volumeauto
```

### Execute

```bash
$ volumeauto -name project_local
```

e.g.

```

$ volumeauto -name project_local
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
            "o": "addr=host.docker.internal,rw",
            "type": "nfs"
        },
        "Scope": "local"
    }
]
```

# License

MIT
