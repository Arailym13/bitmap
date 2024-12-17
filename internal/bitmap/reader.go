package bitmap

import (
	"encoding/binary"
	"errors"
	"os"
)

func ReadBMP(filename string) (*BMPHeader, *DIBHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	bmpHeader := BMPHeader{}
	dibHeader := DIBHeader{}

	if err = binary.Read(file, binary.LittleEndian, &bmpHeader); err != nil {
		return nil, nil, err
	}

	if err = binary.Read(file, binary.LittleEndian, &dibHeader); err != nil {
		return nil, nil, err
	}

	if bmpHeader.FileType[0] != 'B' || bmpHeader.FileType[1] != 'M' {
		return nil, nil, errors.New("invalid BMP format")
	}

	return &bmpHeader, &dibHeader, nil
}
