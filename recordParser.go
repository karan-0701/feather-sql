package main

import (
	"encoding/binary"
	"errors"
	"io"
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

func parseRecord(stream io.Reader, columnCount int) ([]int, error) {
	return nil, nil
}

func parseColumnValue(stream io.Reader, serialType int) ([]byte, int, error) {
	if serialType >= 13 && (serialType%2 == 1) {
		nBytes := int((serialType - 13) / 2)
		nBytesSlice := make([]byte, nBytes)
		_, err := stream.Read(nBytesSlice)
		if err != nil {
			return nil, -1, err
		}
		return nBytesSlice, -1, nil
	} else if serialType == 1 {
		nBytesSerial := make([]byte, 1)
		_, err := stream.Read(nBytesSerial)
		if err != nil {
			return nil, -1, err
		}
		res := int(nBytesSerial[0])
		return nil, res, nil
	} else {

		return nil, -1, errors.New("None of the Serial types exist")
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
