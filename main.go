package main

import (
	"log"

	"github.com/getlantern/systray"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go read()
	systray.Run(onReady, onExit)
}
