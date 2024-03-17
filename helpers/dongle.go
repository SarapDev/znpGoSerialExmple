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

func (d *Dongle) sendCommand(gsp []uint16) []byte {
	gspByte := []byte{}

	for _, item := range gsp {
		gspByte = append(gspByte, byte(item))	
	}

	_, err := d.serial.Write(gspByte)
        if err != nil {
		PrintError("Error while write to port", err)
        }
        
        buf := make([]byte, 128)

	n, err := d.serial.Read(buf)
        if err != nil {
		PrintError("Error while read from port", err)
        }

	return buf[:n]
}

func (d *Dongle) execute(name string, gsp []uint16) {
	log.Println("----------------------------------------------------------------------------------------")
	PrintHexFromUint(fmt.Sprintf("GSP COMMAND %s", name), gsp)

	res := d.sendCommand(gsp)

	PrintHex("DONGLE RESPONSE", res)	

	log.Println("----------------------------------------------------------------------------------------")
}

func (d *Dongle) SysPing() {
	d.execute("SYS_PING", makeCommand(
			[]uint16{0x21, 0x01},
			[]uint16{},
		),
	)
}

func (d *Dongle) SysReset() {
	d.execute("SYS_RESET", makeCommand(
			[]uint16{0x41, 0x00},
			[]uint16{0x00},
		),
	)
}

func (d *Dongle) SysOsalNvWriteDefault() {
	d.execute("SYS_OSAL_NV_WRITE", makeCommand(
			[]uint16{0x21, 0x09},
			[]uint16{0x0087, 0, 1, 2},
		),
	)
}

func (d *Dongle) SysOsalNvWriteEndDevice() {
	d.execute("SYS_OSAL_NV_WRITE", makeCommand(
			[]uint16{0x21, 0x09},
			[]uint16{0, 0x87, 0, 0x01, 0x01},
		),
	)
}

func (d *Dongle) SysOsalNvRead() {
	d.execute("SYS_OSAL_NV_READ", makeCommand(
			[]uint16{0x21, 0x08},
			[]uint16{0x0087, 0},
		),
	)
}

func (d *Dongle) GetDeviceInfo() {
	d.execute("UTIL_GET_DEVICE_INFO", makeCommand(
			[]uint16{0x27, 0x00},
			[]uint16{},
		),
	)
}

func (d *Dongle) EndpointRegister() {
	d.execute("AF_REGISTER", makeCommand(
			[]uint16{0x24, 0x00},
			[]uint16{1, 0x0104, 1, 1, 0, 0, 0},
		),
	)
}

func (d *Dongle) FindingAndBinding() {
	d.execute("APP_CNF_BDB_START_COMMISSIONING", makeCommand(
			[]uint16{0x2F, 0x05},
			[]uint16{0x08},
		),
	)

}
