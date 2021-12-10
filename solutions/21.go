package solutions

import "fmt"

// Task21 - 21.	Реализовать паттерн «адаптер» на любом примере.
func Task21() {
	/*
		The Adapter pattern belongs to structural patterns and is used to structure classes and objects.
		It is used to adapt one system to the requirements of another system.
		The adapter acts as a layer between two objects, converting the calls of one into calls that are understandable to the other.

		Example with USB connector and card reader.
	*/

	usb := &USB{}
	cartReader := &CardReader{}
	client := &Client{}

	client.Connect(usb)                       // connect usb
	adapter := &CartReaderAdapter{cartReader} // use adapter to connect the card reader to usb
	client.Connect(adapter)                   // connect usb

}

type USB struct{}
type CardReader struct{}
type Client struct{}

type CartReaderAdapter struct {
	t *CardReader
}
type Connect interface {
	InsertUSB()
}

func (t *CardReader) InsertCardReader() {
	fmt.Println("CardReader connect successful")
}
func (usb *USB) InsertUSB() {
	fmt.Println("USB connect successful")
}

func (a *CartReaderAdapter) InsertUSB() {
	a.t.InsertCardReader()
}

func (c *Client) Connect(con Connect) {
	con.InsertUSB()
}
