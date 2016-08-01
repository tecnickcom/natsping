package main

import (
	//"fmt"
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestPrefixFieldClashes(t *testing.T) {
	log.WithFields(log.Fields{
		"msg": "additional message",
	}).Info("testing log")
}
