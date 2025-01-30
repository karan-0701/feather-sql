package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	databaseFilePath := args[1]
	command := args[2]

	if command == ".dbinfo" {
		databaseFile, err := os.Open(databaseFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer databaseFile.Close()
		parseHeader, err := parseFrom(databaseFilePath)
		if err != nil {
			fmt.Print(err)
			return
		}

		cellPointers := make([]int, parseHeader.numberOfCells)
		_, err = databaseFile.Seek(100+8, 0)
		if err != nil {
			fmt.Println(err)
			return
		}

		cellPointerReader := make([]byte, parseHeader.numberOfCells*2)
		_, err = databaseFile.Read(cellPointerReader)
		if err != nil {
			fmt.Println(err)
		}

		offset := 0
		for i := 0; i < parseHeader.numberOfCells; i++ {
			cellPointers[i] = int(binary.BigEndian.Uint16(cellPointerReader[offset : offset+2]))
			offset += 2
		}

	}
}
