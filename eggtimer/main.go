package main

import "gioui.org/app"

func main() {
	go func() {
		w := app.NewWindow()

		//_ = w
		for range w.Events() {
		}

	}()

	app.Main()
}
