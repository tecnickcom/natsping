package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func cli() (*cobra.Command, error) {

	// configuration parameters
	cfgParams, err := getConfigParams()
	if err != nil {
		return nil, err
	}

	// application parameters
	appParams := new(params)

	// set the root command
	rootCmd := new(cobra.Command)

	// overwrites the configuration parameters with the ones specified in the command line (if any)
	rootCmd.Flags().StringVarP(&appParams.natsAddress, "natsAddress", "n", cfgParams.natsAddress, "NATS bus Address (nats://ip:port)")
	rootCmd.Flags().StringVarP(&appParams.logLevel, "logLevel", "l", cfgParams.logLevel, "Log level: panic, fatal, error, warning, info, debug")

	rootCmd.Use = "natsping"
	rootCmd.Short = "NATS Bus Ping Command"
	rootCmd.Long = `NATS Bus Ping Command`
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// check config values
		err := checkParams(appParams)
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
