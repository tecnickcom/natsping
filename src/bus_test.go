package main

import (
	"fmt"
	"testing"
)

func TestOpenNatsBusError(t *testing.T) {
	initNatsBus("nats://127.0.0.1:3333")
	err := openNatsBus()
	if err == nil {
		closeNatsBus()
		t.Error(fmt.Errorf("an error was expected"))
	}
}

func TestCheckNatsBusError(t *testing.T) {
	initNatsBus("nats://127.0.0.1:3333")
	err := checkNatsBus()
	if err == nil {
		t.Error(fmt.Errorf("an error was expected"))
	}
}

func TestCheckNatsBus(t *testing.T) {
	initNatsBus("nats://127.0.0.1:4222")
	err := checkNatsBus()
	if err != nil {
		t.Error(fmt.Errorf("an error was not expected"))
	}
}
