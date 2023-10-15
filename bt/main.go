package main

import (
	"time"

	"tinygo.org/x/bluetooth"
)

func main() {
	adapter := bluetooth.DefaultAdapter
	must("enable BLE stack", adapter.Enable())

	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Go Bluetooth",
	}))

	must("start adv", adv.Start())

	println("advertising")
	for {
		time.Sleep(time.Hour)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to" + action + ": " + err.Error())
	}
}
