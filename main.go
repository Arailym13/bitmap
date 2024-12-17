package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"os"
)

type BMPHeader struct {
	FileType   [2]byte
	FileSize   uint32
	Reserved1  uint16
	Reserved2  uint16
	DataOffset uint32
}

type DIBHeader struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XpixelsPerM     uint32
	YpixelsPerM     uint32
	ColorsUsed      uint32
	ImportantColors uint32
}

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

func main() {
	bmpFile := "sample.bmp"
	fmt.Println(bmpFile)
	header := flag.Bool("header", false, "prints header information of bmp file")
	flag.Parse()
	fmt.Println(*header)
	if *header {
		bmpHeader, dibHeader, err := ReadBMP(bmpFile)
		if err != nil {
			fmt.Println("Err:", err)
		}
		fmt.Printf("Сигнатура файла: %c%c\n", bmpHeader.FileType[0], bmpHeader.FileType[1])
		fmt.Printf("Размер файла: %d байт\n", bmpHeader.FileSize)
		fmt.Printf("Смещение до данных изображения: %d байт\n", bmpHeader.DataOffset)
		fmt.Printf("Размер DIB заголовка: %d байт\n", dibHeader.Size)
		fmt.Printf("Ширина изображения: %d пикселей\n", dibHeader.Width)
		fmt.Printf("Высота изображения: %d пикселей\n", dibHeader.Height)
		fmt.Printf("Количество бит на пиксель: %d\n", dibHeader.BitsPerPixel)
	}
}
