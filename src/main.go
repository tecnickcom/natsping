// Copyright (c) 2016-2017 MIRACL UK LTD
// NatsPing for MAAS-SSO
package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	rootCmd, err := cli()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("UNABLE TO START THE PROGRAM")
	}
	// execute the root command and log errors (if any)
	if err = rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("UNABLE TO RUN THE COMMAND")
	}
}
