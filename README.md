# docker-nfs

Auto configuration cli for insert docker nfs driver config into `/etc/exports`.

## Installation

```bash
$ go install github.com/katsew/docker-nfs/cmd/nfsauto
```

## Execute

```bash
$ sudo nfsauto [address - default: localhost]
```

e.g.

```
$ cd /path/to/your/docker/project
$ sudo nfsauto 192.168.33.10
$ cat /etc/exports

# BEGIN - docker-nfs [uid]:[gid]
"/path/to/your/docker/project" 192.168.33.10 -alldirs -mapall=[uid]:[gid]
# END - docker-nfs [uid]:[gid]

```

## License

MIT
