package main

import "github.com/getlantern/systray"

func main() {
	go read()
	systray.Run(onReady, onExit)
}
