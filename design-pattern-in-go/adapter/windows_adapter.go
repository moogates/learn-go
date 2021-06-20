package main

import "fmt"

type WindowsAdapter struct {
    windowMachine *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}
