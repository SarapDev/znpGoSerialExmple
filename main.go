package main

import (
	"serial/helpers"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty.usbserial-140", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
		helpers.PrintError("Error with opening port", err)
        }
       	
	dongle := helpers.NewDongle(s)

	dongle.SysPing()

	dongle.SysReset()

	dongle.SysOsalNvWriteDefault()

	dongle.SysReset()

	dongle.SysOsalNvRead()
}

