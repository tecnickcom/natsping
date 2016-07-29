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


## Getting started (for developers)

This application is written in GO language, please refer to the guides in https://golang.org for getting started.

This project should be cloned in the following directory:
```
$GOPATH/src/github.com/miracl/natsping
```

This project include a Makefile that allows you to test and build the project with simple commands.
To see all available options:
```
make help
```

To build the project
```
make build
```


Alternatively this project can be built inside a Docker container using the command:
```
make dbuild
```

## Running all tests

Before committing the code, please check if it passes all tests using
```
make qa
```

Please check all the available options using
```
make help
```


## Usage

```
Usage:
  natsping [flags]
  natsping [command]

Available Commands:
  version     print this program version

Flags:
  -n, --natsAddress string     NATS bus Address (nats://ip:port) (default "nats://127.0.0.1:4222")
  -l, --logLevel      string  Log level: NONE, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG

Use "natsping [command] --help" for more information about a command.
```

## How it works

The service can be started by issuing the following command (*with the right parameters*):

```
natsping --natsAddress="nats://127.0.0.1:4222 --logLevel=INFO"
```

If no command-line parameters are specified, then the ones in the configuration file (**config.json**) will be used.  
The configuration files can be stored in the current directory or in any of the following (in order of precedence):
* ./
* config/
* $HOME/natsping/
* /etc/natsping/


This service also support secure remote configuration via [Consul](https://www.consul.io/) or [Etcd](https://github.com/coreos/etcd).  
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

This service logs the log messages in the *Stderr* or *Stdout* using the syslog prefixes:


| PREFIX                              | LEVEL | DESCRIPTION                                                                            | OUTPUT |
|:----------------------------------- |:-----:|:-------------------------------------------------------------------------------------- |:------:|
|<nobr> [EMERGENCY] [natsping] </nobr>|   0   |<nobr> **Emergency**: System is unusable                                         </nobr>| Stderr |
|<nobr> [ALERT] [natsping]     </nobr>|   1   |<nobr> **Alert**: Should be corrected immediately                                </nobr>| Stderr |
|<nobr> [CRITICAL] [natsping]  </nobr>|   2   |<nobr> **Critical**: Critical conditions                                         </nobr>| Stderr |
|<nobr> [ERROR] [natsping]     </nobr>|   3   |<nobr> **Error**: Error conditions                                               </nobr>| Stderr |
|<nobr> [WARNING] [natsping]   </nobr>|   4   |<nobr> **Warning**: May indicate that an error will occur if action is not taken </nobr>| Stderr |
|<nobr> [NOTICE] [natsping]    </nobr>|   5   |<nobr> **Notice**: Events that are unusual, but not error conditions             </nobr>| Stderr |
|<nobr> [INFO] [natsping]      </nobr>|   6   |<nobr> **Informational**: Normal operational messages that require no action     </nobr>| Stderr |
|<nobr> [DEBUG] [natsping]     </nobr>|   7   |<nobr> **Debug**: Information useful to developers for debugging the application </nobr>| Stderr |



## Developer(s) Contact

* Nicola Asuni <nicola.asuni@miracl.com>
