// Copyright (c) 2016 MIRACL UK LTD
// NATS Bus Ping Command
package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	rootCmd, err := cli()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to start the command")
	}
	// execute the root command and log errors (if any)
	if err = rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to start the command")
	}
}
