package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats"
)

var busTimeout = time.Duration(BusTimeout) * time.Second
var natsOpts = nats.DefaultOptions
var natsConn *nats.Conn

// initNatsBus initialize the NATS bus options
func initNatsBus(addr string) {
	Log(DEBUG, "initializing NATS bus: '%s'", addr)
	natsOpts.Servers = []string{addr}
}

// openNatsBus connect to the NATS bus
func openNatsBus() (err error) {
	Log(DEBUG, "opening NATS bus connection")
	natsConn, err = natsOpts.Connect()
	if err != nil {
		return fmt.Errorf("Can't connect to the NATS bus: %v", err)
	}
	return nil
}

// closeNatsBus close the NATS bus
func closeNatsBus() {
	Log(DEBUG, "closing NATS bus connection")
	natsConn.Flush()
	natsConn.Close()
}

// checkNatsBus check if the NATS bus is alive
func checkNatsBus() (err error) {
	err = openNatsBus()
	if err == nil {
		closeNatsBus()
	}
	return err
}
