package main

import (
	"time"

	dbus "github.com/godbus/dbus/v5"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := dbus.SessionBus()
	must(err)

	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", "Test", "this is a test", []string{},
		map[string]dbus.Variant{}, int32((5 * time.Second).Milliseconds()))

	must(call.Err)

}
