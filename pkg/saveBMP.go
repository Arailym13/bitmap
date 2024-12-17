package pkg

import (
	"encoding/binary"
	"fmt"
	"image"
	"os"
)

func SaveBMP(filename string, img image.Image) error {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var bmpHeader BMPHeader
	copy(bmpHeader.FileType[:], []byte("BM"))
	fileSize := 14 + 40 + width*height*3
	bmpHeader.FileSize = uint32(fileSize)
	bmpHeader.DataOffset = 54
	dibHeader := DIBHeader{}

	dibHeader.Size = 40
	dibHeader.Width = int32(width)
	dibHeader.Height = int32(height)
	dibHeader.Planes = 1
	dibHeader.BitsPerPixel = 24
	dibHeader.Compression = 0
	dibHeader.ImageSize = uint32(width * height * 3)
	dibHeader.XpixelsPerM = 0
	dibHeader.YpixelsPerM = 0
	dibHeader.ColorsUsed = 0
	dibHeader.ImportantColors = 0

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, bmpHeader); err != nil {
		return fmt.Errorf("failed to write BMP header: %v", err)
	}

	if err := binary.Write(file, binary.LittleEndian, dibHeader); err != nil {
		return fmt.Errorf("failed to write DIB header: %v", err)
	}

	rowSize := (width*3 + 3) & ^3
	buf := make([]byte, rowSize*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			buf[(height-y-1)*rowSize+x*3] = byte(b >> 8)
			buf[(height-y-1)*rowSize+x*3+1] = byte(g >> 8)
			buf[(height-y-1)*rowSize+x*3+2] = byte(r >> 8)
		}
	}

	if _, err := file.Write(buf); err != nil {
		return fmt.Errorf("failed to write pixel data: %v", err)
	}

	return nil
}
