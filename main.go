package main

import (
	"serial/helpers"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty.usbserial-21320", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
		helpers.PrintError("Error with opening port", err)
        }
       	
	dongle := helpers.NewDongle(s)

	dongle.GetDeviceInfo()

	dongle.SysReset()

	dongle.SysOsalNvWriteDefault()

	dongle.SysReset()

	dongle.EndpointRegister()
	
	dongle.FindingAndBinding()
}

