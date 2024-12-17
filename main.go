package main

import (
	"bitmap/internal/flags"
	"bitmap/internal/handler"
	"bitmap/pkg"
	"flag"
	"fmt"
	"image"
	"os"
)

func main() {
	args := os.Args[1:]
	err := handler.InputHandler(args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bmpFile := args[1]
	fmt.Println("Input BMP File:", bmpFile)
	if len(args) > 0 && args[0] == "header" {
		pkg.PrintHeaderInfo(bmpFile)
	}

	mirrorFlag := flag.String("mirror", "", "Mirroring options: horizontal, h, horizontally, hor, vertical, v, vertically, ver")
	flag.Parse()
	fmt.Println(*mirrorFlag)

	if *mirrorFlag != "" {
		mirrorOptions := flags.ParseMirrorFlag(*mirrorFlag)

		file, err := os.Open(bmpFile)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("Error decoding image: %v\n", err)
			os.Exit(1)
		}

		for _, option := range mirrorOptions {
			switch option {
			case "horizontal", "h", "horizontally", "hor":
				img = flags.MirrorHorizontal(img)
			// case "vertical", "v", "vertically", "ver":
			// 	img = mirrorVertical(img)
			default:
				fmt.Println("Error: Unsupported mirror option")
				os.Exit(1)
			}
		}
		outputFile := args[len(args)-1]
		err = pkg.SaveBMP(outputFile, img)
		if err != nil {
			fmt.Printf("Error saving BMP: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Mirrored image saved to: %s\n", outputFile)
	}
}
