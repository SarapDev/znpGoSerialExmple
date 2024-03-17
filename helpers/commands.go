package helpers

// create GSP frame
func makeCommand(command, payload []uint16) []uint16 {
	gsp := []uint16{0xFE}

	var checksum, length uint16

	for range payload {
		length++
	}
	
	command = append(command, payload...)

	checksum ^= length

	for _, item := range command {
		checksum ^= item
	}

	gsp = append(gsp, length)
	gsp = append(gsp, command...)
	gsp = append(gsp, checksum)

	return gsp
}

