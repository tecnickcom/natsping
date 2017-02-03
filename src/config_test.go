package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func getTestCfgParams() *params {
	return &params{
		log: &LogData{
			Level:   "INFO",
			Network: "",
			Address: "",
		},
		natsAddress: "nats://127.0.0.1:4222",
	}
}

func TestCheckParams(t *testing.T) {
	err := checkParams(getTestCfgParams())
	if err != nil {
		t.Error(fmt.Errorf("No errors are expected: %v", err))
	}
}

func TestCheckConfigParametersErrors(t *testing.T) {
	var testCases = []struct {
		fcfg  func(cfg *params) *params
		field string
	}{
		{func(cfg *params) *params { cfg.log.Level = ""; return cfg }, "log.Level"},
		{func(cfg *params) *params { cfg.log.Level = "INVALID"; return cfg }, "log.Level"},
		{func(cfg *params) *params { cfg.natsAddress = ""; return cfg }, "natsAddress"},
	}
	for _, tt := range testCases {
		cfg := getTestCfgParams()
		err := checkParams(tt.fcfg(cfg))
		if err == nil {
			t.Error(fmt.Errorf("An error was expected because the %s field is invalid", tt.field))
		}
	}
}

func TestGetConfigParams(t *testing.T) {
	prm, err := getConfigParams()
	if err != nil {
		t.Error(fmt.Errorf("An error was not expected: %v", err))
	}
	if prm.natsAddress != "nats://127.0.0.1:4222" {
		t.Error(fmt.Errorf("Found different server address than expected, found %s", prm.natsAddress))
	}
	if prm.log.Level != "DEBUG" {
		t.Error(fmt.Errorf("Found different logLevel than expected, found %s", prm.log.Level))
	}
}

func TestGetLocalConfigParams(t *testing.T) {

	// test environment variables
	defer unsetRemoteConfigEnv()
	os.Setenv("NATSPING_REMOTECONFIGPROVIDER", "consul")
	os.Setenv("NATSPING_REMOTECONFIGENDPOINT", "127.0.0.1:98765")
	os.Setenv("NATSPING_REMOTECONFIGPATH", "/config/natsping")
	os.Setenv("NATSPING_REMOTECONFIGSECRETKEYRING", "")

	prm, rprm := getLocalConfigParams()

	if prm.natsAddress != "nats://127.0.0.1:4222" {
		t.Error(fmt.Errorf("Found different server address than expected, found %s", prm.natsAddress))
	}
	if prm.log.Level != "DEBUG" {
		t.Error(fmt.Errorf("Found different logLevel than expected, found %s", prm.log.Level))
	}
	if rprm.remoteConfigProvider != "consul" {
		t.Error(fmt.Errorf("Found different remoteConfigProvider than expected, found %s", rprm.remoteConfigProvider))
	}
	if rprm.remoteConfigEndpoint != "127.0.0.1:98765" {
		t.Error(fmt.Errorf("Found different remoteConfigEndpoint than expected, found %s", rprm.remoteConfigEndpoint))
	}
	if rprm.remoteConfigPath != "/config/natsping" {
		t.Error(fmt.Errorf("Found different remoteConfigPath than expected, found %s", rprm.remoteConfigPath))
	}
	if rprm.remoteConfigSecretKeyring != "" {
		t.Error(fmt.Errorf("Found different remoteConfigSecretKeyring than expected, found %s", rprm.remoteConfigSecretKeyring))
	}

	_, err := getRemoteConfigParams(prm, rprm)
	if err == nil {
		t.Error(fmt.Errorf("A remote configuration error was expected"))
	}

	rprm.remoteConfigSecretKeyring = "/etc/natsping/cfgkey.gpg"
	_, err = getRemoteConfigParams(prm, rprm)
	if err == nil {
		t.Error(fmt.Errorf("A remote configuration error was expected"))
	}
}

// Test real Consul provider
// To activate this define the environmental variable NATSPING_LIVECONSUL
func TestGetConfigParamsRemote(t *testing.T) {

	enable := os.Getenv("NATSPING_LIVECONSUL")
	if enable == "" {
		return
	}

	// test environment variables
	defer unsetRemoteConfigEnv()
	os.Setenv("NATSPING_REMOTECONFIGPROVIDER", "consul")
	os.Setenv("NATSPING_REMOTECONFIGENDPOINT", "127.0.0.1:8500")
	os.Setenv("NATSPING_REMOTECONFIGPATH", "/config/natsping")
	os.Setenv("NATSPING_REMOTECONFIGSECRETKEYRING", "")

	// load a specific config file just for testing
	oldCfg := ConfigPath
	viper.Reset()
	for k := range ConfigPath {
		ConfigPath[k] = "wrong/path/"
	}
	defer func() { ConfigPath = oldCfg }()

	prm, err := getConfigParams()
	if err != nil {
		t.Error(fmt.Errorf("An error was not expected: %v", err))
	}
	if prm.natsAddress != "nats://127.0.0.1:4222" {
		t.Error(fmt.Errorf("Found different server address than expected, found %s", prm.natsAddress))
	}
	if prm.log.Level != "debug" {
		t.Error(fmt.Errorf("Found different logLevel than expected, found %s", prm.log.Level))
	}
}

func TestCliWrongConfigError(t *testing.T) {

	// test environment variables
	defer unsetRemoteConfigEnv()
	os.Setenv("NATSPING_REMOTECONFIGPROVIDER", "consul")
	os.Setenv("NATSPING_REMOTECONFIGENDPOINT", "127.0.0.1:999999")
	os.Setenv("NATSPING_REMOTECONFIGPATH", "/config/wrong")
	os.Setenv("NATSPING_REMOTECONFIGSECRETKEYRING", "")

	// load a specific config file just for testing
	oldCfg := ConfigPath
	viper.Reset()
	for k := range ConfigPath {
		ConfigPath[k] = "wrong/path/"
	}
	defer func() { ConfigPath = oldCfg }()

	_, err := cli()
	if err == nil {
		t.Error(fmt.Errorf("An error was expected"))
		return
	}
}

// unsetRemoteConfigEnv clear the environmental variables used to set the remote configuration
func unsetRemoteConfigEnv() {
	os.Setenv("NATSPING_REMOTECONFIGPROVIDER", "")
	os.Setenv("NATSPING_REMOTECONFIGENDPOINT", "")
	os.Setenv("NATSPING_REMOTECONFIGPATH", "")
	os.Setenv("NATSPING_REMOTECONFIGSECRETKEYRING", "")
}
