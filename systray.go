package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	trayIcon, err := Asset("data/vcard.ico")
	if err != nil {
		log.Println("vcard.ico not available. Using default one.")
		trayIcon = icon.Data
	}

	systray.SetIcon(trayIcon)
	systray.SetTitle("golksd")
	systray.SetTooltip("golksd - identity card reader")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit application")

	go func() {
		<-mQuitOrig.ClickedCh
		log.Print("Requesting systray quit...")
		systray.Quit()
		log.Println("done!")
	}()
}

func onExit() {
	log.Println("Exiting...")
	os.Exit(0)
}
