package main

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// params struct contains the application parameters
type remoteConfigParams struct {
	remoteConfigProvider      string // remote configuration source ("consul", "etcd")
	remoteConfigEndpoint      string // remote configuration URL (ip:port)
	remoteConfigPath          string // remote configuration path where to search fo the configuration file ("/config/natsping")
	remoteConfigSecretKeyring string // path to the openpgp secret keyring used to decript the remote configuration data ("/etc/natsping/configkey.gpg")
}

// isEmpty returns true if all the fields are empty strings
func (rcfg remoteConfigParams) isEmpty() bool {
	return rcfg.remoteConfigProvider == "" && rcfg.remoteConfigEndpoint == "" && rcfg.remoteConfigPath == "" && rcfg.remoteConfigSecretKeyring == ""
}

// params struct contains the application parameters
type params struct {
	log         *LogData // Log level: EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG.
	natsAddress string   // NATS bus Address (nats://ip:port)
}

var configDir string
var appParams = &params{}

// getConfigParams returns the configuration parameters
func getConfigParams() (par params, err error) {
	cfg, rcfg, err := getLocalConfigParams()
	if err != nil {
		return par, err
	}
	return getRemoteConfigParams(cfg, rcfg)
}

// getLocalConfigParams returns the local configuration parameters
func getLocalConfigParams() (cfg params, rcfg remoteConfigParams, err error) {

	viper.Reset()

	// set default remote configuration values
	viper.SetDefault("remoteConfigProvider", RemoteConfigProvider)
	viper.SetDefault("remoteConfigEndpoint", RemoteConfigEndpoint)
	viper.SetDefault("remoteConfigPath", RemoteConfigPath)
	viper.SetDefault("remoteConfigSecretKeyring", RemoteConfigSecretKeyring)

	// set default configuration values

	viper.SetDefault("log.level", LogLevel)
	viper.SetDefault("log.network", LogNetwork)
	viper.SetDefault("log.address", LogAddress)

	viper.SetDefault("natsAddress", NatsAddress)

	// name of the configuration file without extension
	viper.SetConfigName("config")

	// configuration type
	viper.SetConfigType("json")

	if configDir != "" {
		viper.AddConfigPath(configDir)
	}

	// add local configuration paths
	for _, cpath := range ConfigPath {
		viper.AddConfigPath(cpath)
	}

	// Find and read the local configuration file (if any)
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, rcfg, err
	}

	// read configuration parameters
	cfg = getViperParams()

	// support environment variables for the remote configuration
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(ProgramName, "-", "_", -1)) // will be uppercased automatically
	envVar := []string{
		"remoteConfigProvider",
		"remoteConfigEndpoint",
		"remoteConfigPath",
		"remoteConfigSecretKeyring",
	}
	for _, ev := range envVar {
		err = viper.BindEnv(ev)
		if err != nil {
			return cfg, rcfg, err
		}
	}

	rcfg = remoteConfigParams{
		remoteConfigProvider:      viper.GetString("remoteConfigProvider"),
		remoteConfigEndpoint:      viper.GetString("remoteConfigEndpoint"),
		remoteConfigPath:          viper.GetString("remoteConfigPath"),
		remoteConfigSecretKeyring: viper.GetString("remoteConfigSecretKeyring"),
	}

	return cfg, rcfg, nil
}

// getRemoteConfigParams returns the remote configuration parameters
func getRemoteConfigParams(cfg params, rcfg remoteConfigParams) (params, error) {

	if rcfg.isEmpty() {
		return cfg, nil
	}

	viper.Reset()

	// set default configuration values

	viper.SetDefault("log.level", cfg.log.Level)
	viper.SetDefault("log.network", cfg.log.Network)
	viper.SetDefault("log.address", cfg.log.Address)

	viper.SetDefault("natsAddress", cfg.natsAddress)

	// configuration type
	viper.SetConfigType("json")

	// add remote configuration provider
	var err error
	if rcfg.remoteConfigSecretKeyring == "" {
		err = viper.AddRemoteProvider(rcfg.remoteConfigProvider, rcfg.remoteConfigEndpoint, rcfg.remoteConfigPath)
	} else {
		err = viper.AddSecureRemoteProvider(rcfg.remoteConfigProvider, rcfg.remoteConfigEndpoint, rcfg.remoteConfigPath, rcfg.remoteConfigSecretKeyring)
	}
	if err == nil {
		// try to read the remote configuration (if any)
		err = viper.ReadRemoteConfig()
	}
	if err != nil {
		return cfg, err
	}

	// read configuration parameters
	return getViperParams(), nil
}

// getViperParams reads the config params via Viper
func getViperParams() params {
	return params{

		log: &LogData{
			Level:   viper.GetString("log.level"),
			Network: viper.GetString("log.network"),
			Address: viper.GetString("log.address"),
		},

		natsAddress: viper.GetString("natsAddress"),
	}
}

// checkParams cheks if the configuration parameters are valid
func checkParams(prm *params) error {
	// Log
	if prm.log.Level == "" {
		return errors.New("log Level is empty")
	}
	err := prm.log.setLog()
	if err != nil {
		return err
	}

	if prm.natsAddress == "" {
		return errors.New("natsAddress is empty")
	}

	return nil
}
