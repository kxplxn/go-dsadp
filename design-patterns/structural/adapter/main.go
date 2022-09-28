package main

import "fmt"

// Client represents an existing business logic in the program that we want to
// extend.
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("client inserts lightning connector into computer")
	com.InsertIntoLightningPort()
}

// Computer is the Client Interface. It describes a protocol that other classes
// must follow to be able to collaborate with the client code.
type Computer interface {
	InsertIntoLightningPort()
}

// Mac is an existing Service that is compatible with the Client logic.
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("lightning connector is plugged into mac machine")
}

// Windows is a new Service that the client cannot use directly because it has
// an incompatible interface.
type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("usb connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("adapter converts lightning signal to usb")
	w.windowsMachine.InsertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	win := &Windows{}
	winAdapter := &WindowsAdapter{
		windowsMachine: win,
	}
	client.InsertLightningConnectorIntoComputer(winAdapter)
}
