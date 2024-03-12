package helpers

import (
	"strings"
	"log"
	"encoding/hex"
)

// print hex bytes as string separeted each 2 bytes
func PrintHex(tag string,hexSlice []byte) {
	var final string	
	hexString := hex.EncodeToString(hexSlice)

	for k, char := range hexString {
		if k % 2 == 0 {
			final += " 0x"
		}

		final += strings.ToUpper(string(char))
	}

	log.Printf("%s: %s", tag, strings.Trim(final, " "))
}

func PrintError(customMessage string, err error) {
	log.Fatalf("%s : %s", customMessage, err)
}

