package utils

import (
	"bitmap/internal/bitmap"
	"fmt"
)

func PrintHeaderInfo(bmpHeader *bitmap.BMPHeader, dibHeader *bitmap.DIBHeader) {
	fmt.Println("-----BMP Header-----")
	fmt.Printf("File Type: %c%c\n", bmpHeader.FileType[0], bmpHeader.FileType[1])
	fmt.Printf("File Size: %d bytes\n", bmpHeader.FileSize)
	fmt.Printf("Offset to image data: %d bytes\n", bmpHeader.DataOffset)
	fmt.Println("-----DIB Header-----")
	fmt.Printf("DIB Header Size: %d bytes\n", dibHeader.Size)
	fmt.Printf("Image width: %d pixels\n", dibHeader.Width)
	fmt.Printf("Image height: %d pixels\n", dibHeader.Height)
	fmt.Printf("Bit counts per pixel: %d\n", dibHeader.BitsPerPixel)
	fmt.Printf("Image Size: %d bytes\n", dibHeader.ImageSize)
}
