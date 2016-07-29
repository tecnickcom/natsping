// Copyright (c) 2016 MIRACL UK LTD
// NATS Bus Ping Command
package main

func main() {
	rootCmd, err := cli()
	if err != nil {
		Log(CRITICAL, "unable to start the command: %v", err)
	}
	// execute the root command and log errors (if any)
	if err = rootCmd.Execute(); err != nil {
		Log(CRITICAL, "unable to start the command: %v", err)
	}
}
