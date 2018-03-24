package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	trayIcon, err := Asset("data/vcard.ico")
	if err != nil {
		fmt.Println("vcard.ico not available. Using default one.")
		trayIcon = icon.Data
	}

	systray.SetIcon(trayIcon)
	systray.SetTitle("golksd")
	systray.SetTooltip("golksd - identity card reader")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit application")

	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Print("Requesting systray quit...")
		systray.Quit()
		fmt.Println("done!")
	}()
}

func onExit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}
