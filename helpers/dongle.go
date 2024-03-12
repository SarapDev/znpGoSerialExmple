package helpers

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

type Dongle struct {
	serial *serial.Port 
}

func NewDongle(serial *serial.Port) *Dongle {
	return &Dongle{
		serial: serial,
	}
}

func (d *Dongle) sendCommand(gsp []byte) []byte {
	_, err := d.serial.Write(gsp)
        if err != nil {
		PrintError("Error while write to port", err)
        }
        
        buf := make([]byte, 32)

	n, err := d.serial.Read(buf)
        if err != nil {
		PrintError("Error while read from port", err)
        }

	return buf[:n]
}

func (d *Dongle) execute(name string, gsp []byte) {
	log.Println("----------------------------------------------------------------------------------------")

	PrintHex(fmt.Sprintf("GSP COMMAND %s", name), gsp)

	res := d.sendCommand(gsp)

	PrintHex("DONGLE RESPONSE", res)	

	log.Println("----------------------------------------------------------------------------------------")
}

func (d *Dongle) SysPing() {
	d.execute("SYS_PING", makeCommand(
			[]byte{0x21, 0x01},
			[]byte{},
		),
	)
}

func (d *Dongle) SysReset() {
	d.execute("SYS_RESET", makeCommand(
			[]byte{0x41, 0x00},
			[]byte{0x00},
		),
	)
}

func (d *Dongle) SysOsalNvWriteDefault() {
	d.execute("SYS_OSAL_NV_WRITE", makeCommand(
			[]byte{0x21, 0x09},
			[]byte{0x87, 0, 0x01, 0x02},
		),
	)
}

func (d *Dongle) SysOsalNvRead() {
	d.execute("SYS_OSAL_NV_READ", makeCommand(
			[]byte{0x21, 0x08},
			[]byte{0x87, 0},
		),
	)
}

