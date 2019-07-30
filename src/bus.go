package main

import (
	"fmt"

	"github.com/nats-io/nats"
	log "github.com/sirupsen/logrus"
)

var natsOpts = nats.DefaultOptions
var natsConn *nats.Conn

// initNatsBus initialize the NATS bus options
func initNatsBus(addr string) {
	log.WithFields(log.Fields{
		"nats": addr,
	}).Debug("initializing NATS bus")
	natsOpts.Servers = []string{addr}
}

// openNatsBus connect to the NATS bus
func openNatsBus() (err error) {
	log.WithFields(log.Fields{
		"nats": natsOpts.Servers,
	}).Debug("opening NATS bus connection")
	natsConn, err = natsOpts.Connect()
	if err != nil {
		return fmt.Errorf("can't connect to the NATS bus: %v", err)
	}
	return nil
}

// closeNatsBus close the NATS bus
func closeNatsBus() {
	log.WithFields(log.Fields{
		"nats": natsOpts.Servers,
	}).Debug("closing NATS bus connection")
	err := natsConn.Flush()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("flush NATS bus connection")
	}
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
