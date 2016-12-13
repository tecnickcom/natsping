# natsping

*NATS Bus Test Component*

[![Master Branch](https://img.shields.io/badge/-master:-gray.svg)](https://github.com/miracl/natsping/tree/master)
[![Master Build Status](https://secure.travis-ci.org/miracl/natsping.png?branch=master)](https://travis-ci.org/miracl/natsping?branch=master)
[![Master Coverage Status](https://coveralls.io/repos/miracl/natsping/badge.svg?branch=master&service=github)](https://coveralls.io/github/miracl/natsping?branch=master)

[![Develop Branch](https://img.shields.io/badge/-develop:-gray.svg)](https://github.com/miracl/natsping/tree/develop)
[![Develop Build Status](https://secure.travis-ci.org/miracl/natsping.png?branch=develop)](https://travis-ci.org/miracl/natsping?branch=develop)
[![Develop Coverage Status](https://coveralls.io/repos/miracl/natsping/badge.svg?branch=develop&service=github)](https://coveralls.io/github/miracl/natsping?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/miracl/natsping)](https://goreportcard.com/report/github.com/miracl/natsping)

* **category**:    Tool
* **author**:      Nicola Asuni <nicola.asuni@miracl.com>
* **copyright**:   2016 MIRACL UK LTD
* **license**:     ASL 2.0 - http://www.apache.org/licenses/LICENSE-2.0
* **link**:        https://github.com/miracl/natsping

## Description

*NATS Bus Ping Command.*

This command-line program allows to ping a [NATS](http://nats.io) bus to see if it is alive.


## Quick Start

This project includes a Makefile that allows you to test and build the project in a Linux-compatible system with simple commands.  
All the artifacts and reports produced using this Makefile are stored in the *target* folder.  

All the packages listed in the *resources/DockerDev/Dockerfile* file are required in order to build and test all the library options in the current environment. Alternatively, everything can be built inside a [Docker](https://www.docker.com) container using the command "make dbuild".

To see all available options:
```
make help
```

To build the project inside a Docker container (requires Docker):
```
make dbuild
```

To build a particular set of options inside a Docker container:
```
MAKETARGET='buildall' make dbuild
```
The list of pre-defined options can be listed by typing ```make```


The base Docker building environment is defined in the following Dockerfile:
```
resources/DockerDev/Dockerfile
```

To execute all the default test builds and generate reports in the current environment:
```
make qa
```

To format the code (please use this command before submitting any pull request):
```
make format
```

## Useful Docker commands

To manually create the container you can execute:
```
docker build --tag="miracl/natspingdev" .
```

To log into the newly created container:
```
docker run -t -i miracl/natspingdev /bin/bash
```

To get the container ID:
```
CONTAINER_ID=`docker ps -a | grep miracl/natspingdev | cut -c1-12`
```

To delete the newly created docker container:
```
docker rm -f $CONTAINER_ID
```

To delete the docker image:
```
docker rmi -f miracl/natspingdev
```

To delete all containers
```
docker rm $(docker ps -a -q)
```

To delete all images
```
docker rmi $(docker images -q)
```


## Usage

```
Usage:
  natsping [flags]
  natsping [command]

Available Commands:
  version     print this program version

Flags:
  -c, --configDir   string  Configuration directory to be added on top of the search list
  -l, --logLevel     string  Log level: EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG
  -n, --natsAddress  string  NATS bus Address (nats://ip:port) (default "nats://127.0.0.1:4222")

Use "natsping [command] --help" for more information about a command.
```

## How it works

The program can be started by issuing the following command (*with the right parameters*):

```
natsping --natsAddress="nats://127.0.0.1:4222 --logLevel=INFO"
```

If no command-line parameters are specified, then the ones in the configuration file (**config.json**) will be used.  
The configuration files can be stored in the current directory or in any of the following (in order of precedence):
* ./
* config/
* $HOME/natsping/
* /etc/natsping/


This program also support secure remote configuration via [Consul](https://www.consul.io/) or [Etcd](https://github.com/coreos/etcd).  
The remote configuration server can be defined either in the local configuration file using the following parameters, or with environment variables:

* **remoteConfigProvider** : remote configuration source ("consul", "etcd");
* **remoteConfigEndpoint** : remote configuration URL (ip:port);
* **remoteConfigPath** : remote configuration path where to search fo the configuration file (e.g. "/config/natsping");
* **remoteConfigSecretKeyring** : path to the openpgp secret keyring used to decript the remote configuration data (e.g. "/etc/natsping/configkey.gpg"); if empty a non secure connection will be used instead;

The equivalent environment variables are:

* NATSPING_REMOTECONFIGPROVIDER
* NATSPING_REMOTECONFIGENDPOINT
* NATSPING_REMOTECONFIGPATH
* NATSPING_REMOTECONFIGSECRETKEYRING


The natsping command exit with the status 0 if the NATS bus is responding, otherwise it generates an error log message and exit with 1.


## Logs

This program logs the log messages in json format.
For example:
```
{"datetime":"2016-10-06T15:00:39Z","hostname":"nickxm","level":"debug","msg":"initializing NATS bus","nats":"nats://127.0.0.1:3333","program":"natsping","release":"1","timestamp":1475766039219189166,"version":"1.2.1"}

```

## Developer(s) Contact

* Nicola Asuni <nicola.asuni@miracl.com>
