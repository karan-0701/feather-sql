package main

import (
	"encoding/binary"
	"os"
)

type PageHeader struct {
	pageType              int
	firstFreeBlockStart   int
	numberOfCells         int
	startOfTheContentArea int
	fragmentedFreeBytes   int
}

func newPageHeader(pageType int, firstFreeBlockStart int, numberofCells int, startOfContentArea int, fragmentedFreeBytes int) *PageHeader {
	return &PageHeader{
		pageType,
		firstFreeBlockStart,
		numberofCells,
		startOfContentArea,
		fragmentedFreeBytes,
	}
}

func parseFrom(databaseFilePath string) (*PageHeader, error) {

	file, err := os.Open(databaseFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// move the pointer by 100 bytes to skip the database header
	_, err = file.Seek(100, 0)
	if err != nil {
		return nil, err
	}

	// read the B-Tree page header that is 8 bytes long
	header := make([]byte, 8)
	_, err = file.Read(header)
	if err != nil {
		return nil, err
	}

	pageType := int(header[0])
	firstFreeBlockStart := int(binary.BigEndian.Uint16(header[1:3]))
	numberOfCells := int(binary.BigEndian.Uint16(header[3:5]))
	startOfContentArea := int(binary.BigEndian.Uint16(header[5:7]))
	fragmentedFreeBytes := int(header[7])

	return newPageHeader(pageType, firstFreeBlockStart, numberOfCells, startOfContentArea, fragmentedFreeBytes), nil
}
