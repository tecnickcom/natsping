package main

// ProgramName contains this program name
const ProgramName = "natsping"

// ProgramVersion contains this program version
// This is automatically populated by the Makefile using the value from the VERSION file
var ProgramVersion = "0.0.0"

// ProgramRelease contains this program release number (or build number)
// This is automatically populated by the Makefile using the value from the RELEASE file
var ProgramRelease = "0"

// NatsAddress is the default NATS bus address
const NatsAddress = "nats://127.0.0.1:4222"

// BusTimeout is the default NATS bus connection timeout in seconds
const BusTimeout = 1

// ConfigPath list the local paths where to look for configuration files (in order)
var ConfigPath = [...]string{
	"../resources/test/etc/" + ProgramName + "/",
	"./",
	"config/",
	"$HOME/." + ProgramName + "/",
	"/etc/" + ProgramName + "/",
}

// LogLevel defines the default log level: NONE, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG
const LogLevel = "info"

// RemoteConfigProvider is the remote configuration source ("consul", "etcd")
const RemoteConfigProvider = ""

// RemoteConfigEndpoint is the remote configuration URL (ip:port)
const RemoteConfigEndpoint = ""

// RemoteConfigPath is the remote configuration path where to search fo the configuration file ("/config/natsping")
const RemoteConfigPath = ""

// RemoteConfigSecretKeyring is the path to the openpgp secret keyring used to decript the remote configuration data ("/etc/natsping/configkey.gpg")
const RemoteConfigSecretKeyring = ""
