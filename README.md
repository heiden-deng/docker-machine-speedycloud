# Docker Machine speedycloud Driver

This is a plugin for [Docker Machine](https://docs.docker.com/machine/) allowing
to create Docker hosts locally on SpeedyCloud (http://www.speedycloud.cn/)

## Requirements
* OS X 10.9+
* [Docker Machine](https://docs.docker.com/machine/) 0.5.1+ (is bundled to
  [Docker Toolbox](https://www.docker.com/docker-toolbox) 1.9.1+)
* [speedycloud Desktop](http://www.speedycloud.com/products/desktop/) 11.0.0+ **Pro** or
**Business** edition (_Standard edition is not supported!_)

## Installation
Install via Homebrew:

```console
$ brew install docker-machine-speedycloud
```

To install this plugin manually, download the binary `docker-machine-driver-speedycloud`
and  make it available by `$PATH`, for example by putting it to `/usr/local/bin/`:

```console
$ curl -L https://github.com/speedycloud/docker-machine-speedycloud/releases/download/v1.2.2/docker-machine-driver-speedycloud > /usr/local/bin/docker-machine-driver-speedycloud

$ chmod +x /usr/local/bin/docker-machine-driver-speedycloud
```

The latest version of `docker-machine-driver-speedycloud` binary is available on
the ["Releases"](https://github.com/speedycloud/docker-machine-speedycloud/releases) page.

## Usage
Official documentation for Docker Machine [is available here](https://docs.docker.com/machine/).

To create a speedycloud Desktop virtual machine for Docker purposes just run this
command:

```
$ docker-machine create --driver=speedycloud prl-dev
```

Available options:

 - `--speedycloud-boot2docker-url`: The URL of the boot2docker image.
 - `--speedycloud-disk-size`: Size of disk for the host VM (in MB).
 - `--speedycloud-memory`: Size of memory for the host VM (in MB).
 - `--speedycloud-cpu-count`: Number of CPUs to use to create the VM (-1 to use the number of CPUs available).
 - `--speedycloud-no-share`: Disable the sharing of `/Users` directory

The `--speedycloud-boot2docker-url` flag takes a few different forms. By
default, if no value is specified for this flag, Machine will check locally for
a boot2docker ISO. If one is found, that will be used as the ISO for the
created machine. If one is not found, the latest ISO release available on
[boot2docker/boot2docker](https://github.com/boot2docker/boot2docker) will be
downloaded and stored locally for future use. Note that this means you must run
`docker-machine upgrade` deliberately on a machine if you wish to update the "cached"
boot2docker ISO.

This is the default behavior (when `--speedycloud-boot2docker-url=""`), but the
option also supports specifying ISOs by the `http://` and `file://` protocols.

Environment variables and default values:

| CLI option                    | Environment variable        | Default                  |
|-------------------------------|-----------------------------|--------------------------|
| `--speedycloud-boot2docker-url` | `speedycloud_BOOT2DOCKER_URL` | *Latest boot2docker url* |
| `--speedycloud-cpu-count`       | `speedycloud_CPU_COUNT`       | `1`                      |
| `--speedycloud-disk-size`       | `speedycloud_DISK_SIZE`       | `20000`                  |
| `--speedycloud-memory`          | `speedycloud_MEMORY_SIZE`     | `1024`                   |
| `--speedycloud-no-share`        | -                           | `false`                  |

## Development

### Build from Source
If you wish to work on speedycloud Driver for Docker machine, you'll first need
[Go](http://www.golang.org) installed (version 1.7+ is required).
Make sure Go is properly installed, including setting up a [GOPATH](http://golang.org/doc/code.html#GOPATH).

Run these commands to build the plugin binary:

```bash
$ go get -d github.com/speedycloud/docker-machine-speedycloud
$ cd $GOPATH/github.com/speedycloud/docker-machine-speedycloud
$ make build
```

After the build is complete, `bin/docker-machine-driver-speedycloud` binary will
be created. If you want to copy it to the `${GOPATH}/bin/`, run `make install`.

### Acceptance Tests

We use [BATS](https://github.com/sstephenson/bats) for acceptance testing, so,
[install it](https://github.com/sstephenson/bats#installing-bats-from-source) first.

You also need to build the plugin binary by calling `make build`.

Then you can run acceptance tests using this command:

```bash
$ make test-acceptance
```

Acceptance tests will invoke the general `docker-machine` binary available by
`$PATH`. If you want to specify it explicitly, just set `MACHINE_BINARY` env variable:

```bash
$ MACHINE_BINARY=/path/to/docker-machine make test-acceptance
```

## Authors

* Mikhail Zholobov ([@legal90](https://github.com/legal90))
* Rickard von Essen ([@rickard-von-essen](https://github.com/rickard-von-essen))
