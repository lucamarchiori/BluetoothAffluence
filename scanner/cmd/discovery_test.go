package main

import (
	"reflect"
	"testing"
)

func TestAddDevice(t *testing.T) {
	var gotten []Device
	var wanted []Device
	addDevice(&gotten, Device{Address: "aa"})
	addDevice(&gotten, Device{Address: "bb"})
	addDevice(&gotten, Device{Address: "bb"})
	addDevice(&gotten, Device{Address: "bb"})
	addDevice(&gotten, Device{Address: "cc"})
	addDevice(&gotten, Device{Address: "cc"})

	wanted = []Device{{Address: "aa"}, {Address: "bb"}, {Address: "cc"}}

	if !reflect.DeepEqual(gotten, wanted) {
		log.Info("Error: wanted {%v}, gotten {%v}", wanted, gotten)
	}
}
