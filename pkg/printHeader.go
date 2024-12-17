package pkg

import "fmt"

func PrintHeaderInfo(filename string) {
	bmpHeader, dibHeader, err := ReadBMP(filename)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Printf("File Signature: %c%c\n", bmpHeader.FileType[0], bmpHeader.FileType[1])
	fmt.Printf("File Size: %d байт\n", bmpHeader.FileSize)
	fmt.Printf("Offset to image data: %d bytes\n", bmpHeader.DataOffset)
	fmt.Printf("DIB Header Size: %d bytes\n", dibHeader.Size)
	fmt.Printf("Image width: %d pixels\n", dibHeader.Width)
	fmt.Printf("Image height: %d pixels\n", dibHeader.Height)
	fmt.Printf("Bit counts per pixel: %d\n", dibHeader.BitsPerPixel)
}
