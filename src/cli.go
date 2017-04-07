package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cli() (*cobra.Command, error) {

	var logLevel string
	var natsAddress string

	rootCmd := new(cobra.Command)
	rootCmd.Flags().StringVarP(&configDir, "configDir", "c", "", "Configuration directory to be added on top of the search list")
	rootCmd.Flags().StringVarP(&logLevel, "logLevel", "o", "*", "Log level: EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG")
	rootCmd.Flags().StringVarP(&natsAddress, "natsAddress", "n", "*", "NATS bus Address (nats://ip:port)")
	err := rootCmd.ParseFlags(os.Args)
	if err != nil {
		return nil, err
	}

	rootCmd.Use = "natsping"
	rootCmd.Short = "NATS Bus Ping Command"
	rootCmd.Long = `NATS Bus Ping Command`
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {

		// configuration parameters
		cfgParams, err := getConfigParams()
		if err != nil {
			return err
		}
		appParams = &cfgParams
		if logLevel != "*" {
			appParams.log.Level = logLevel
		}
		if natsAddress != "*" {
			appParams.natsAddress = natsAddress
		}

		// check values
		err = checkParams(appParams)
		if err != nil {
			return err
		}

		// initialize the NATS bus
		initNatsBus(appParams.natsAddress)

		// check if the bus is alive
		return checkNatsBus()
	}

	// sub-command to print the version
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "print this program version",
		Long:  `print this program version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(ProgramVersion)
		},
	}
	rootCmd.AddCommand(versionCmd)

	return rootCmd, nil
}
