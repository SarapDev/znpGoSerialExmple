package helpers

import (
	"encoding/hex"
	"log"
	"strings"
)

// print hex bytes as string separeted each 2 bytes
func PrintHexFromUint(tag string,hexSlice []uint16) {
	hexSliceByte := []byte{}

	for _, item := range hexSlice {
		hexSliceByte = append(hexSliceByte, byte(item))
	}

	PrintHex(tag, hexSliceByte)
}

func PrintHex(tag string,hexSlice []byte) {
	var final string

	hexString := hex.EncodeToString(hexSlice)
	for k, char := range hexString {
		if k % 2 == 0 {
			final += " 0x"
		}

		final += strings.ToUpper(string(char))
	}

	log.Println(tag, final)
}

func PrintError(customMessage string, err error) {
	log.Fatalf("%s : %s", customMessage, err)
}

