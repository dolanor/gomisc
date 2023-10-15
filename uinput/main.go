package main

import "github.com/sashko/go-uinput"

func main() {
	kbd, err := uinput.CreateKeyboard()
	if err != nil {
		panic(err)
	}
	defer kbd.Close()

	kbd.KeyDown(uinput.KeyLeftShift)
	kbd.KeyPress(uinput.KeyG)
	kbd.KeyUp(uinput.KeyLeftShift)

	kbd.KeyPress(uinput.KeyO)
}
