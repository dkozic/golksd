package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dkozic/golksd/sd"
	"github.com/ebfe/scard"
)

func die(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func waitUntilCardPresent(ctx *scard.Context, readers []string, msg chan int) {
	rs := make([]scard.ReaderState, len(readers))
	for i := range rs {
		rs[i].Reader = readers[i]
		rs[i].CurrentState = scard.StateUnaware
	}

	for {
		for i := range rs {
			fmt.Printf("%d %#v\n", i, rs[i])
			if rs[i].CurrentState != rs[i].EventState {
				if (rs[i].EventState&scard.StatePresent != 0) && (rs[i].CurrentState&scard.StatePresent != rs[i].EventState&scard.StatePresent) && (rs[i].EventState&scard.StateMute == 0) {
					fmt.Printf("send %d\n", i)
					msg <- i
				}
				rs[i].CurrentState = rs[i].EventState
			}
		}
		fmt.Println("Before GetStatusChange")
		err := ctx.GetStatusChange(rs, -1)
		fmt.Println("After GetStatusChange")

		if err != nil {
			fmt.Println("Error GetStatusChange: ", err)
			time.Sleep(2 * time.Second)
		}
	}
}

func read() {

	// Establish a PC/SC context
	fmt.Print("Calling scard.EstablishContext...")
	context, err := scard.EstablishContext()
	if err != nil {
		fmt.Println("Error EstablishContext:", err)
		return
	}
	fmt.Println("done.")

	// Release the PC/SC context (when needed)
	defer context.Release()

	// List available readers
	fmt.Print("Calling context.ListReaders...")
	readers, err := context.ListReaders()
	if err != nil {
		fmt.Println("Error ListReaders:", err)
		return
	}
	fmt.Println("done.")

	fmt.Printf("Found %d readers:\n", len(readers))
	for i, reader := range readers {
		fmt.Printf("[%d] %s\n", i, reader)
	}

	if len(readers) > 0 {
		msg := make(chan int)

		fmt.Println("Waiting for a Card...")
		go waitUntilCardPresent(context, readers, msg)

		for {
			index := <-msg
			fmt.Printf("receive %d\n", index)

			// Use the first reader
			reader := readers[index]
			fmt.Println("Using reader:", reader)

			// Connect to the card
			card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
			if err != nil {
				fmt.Println("Error Connect:", err)
				return
			}

			// Disconnect (when needed)
			defer card.Disconnect(scard.LeaveCard)

			readCard(card)
		}
	}
}

func readCard(card *scard.Card) {
	data, err := sd.ReadEsd(card)
	if err != nil {
		fmt.Printf("Error %#v\n", err)
	}

	log.Printf("EsdData %#v\n", data)

}
func readCard1(card *scard.Card) {

	fmt.Println("Card status:")
	status, err := card.Status()
	if err != nil {
		die(err)
	}

	fmt.Printf("\treader: %s\n\tstate: %x\n\tactive protocol: %x\n\tatr: % x\n",
		status.Reader, status.State, status.ActiveProtocol, status.Atr)

	// Send select APDU
	var cmd_select = []byte{0x00, 0xa4, 0x04, 0x00, 0x0A, 0xA0,
		0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0x0C, 0x06, 0x01}
	rsp, err := card.Transmit(cmd_select)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println(rsp)

	// Send command APDU
	var cmd_command = []byte{0x00, 0x00, 0x00, 0x00}
	rsp, err = card.Transmit(cmd_command)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println(rsp)
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println()
}
