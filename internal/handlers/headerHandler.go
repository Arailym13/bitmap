package handlers

import (
	"bitmap/internal/bitmap"
	"bitmap/utils"
	"fmt"
	"os"
	"strings"
)

func HandleHeader(args []string) {
	parseHeaderArgs(args)
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: 'header' requires exactly one argument (source file)")
		os.Exit(1)
	}
	filename := args[1]
	fmt.Println(filename)
	if !utils.IsValidBMPFile(filename) {
		fmt.Fprintln(os.Stderr, "Error: invalid BMP file format")
		os.Exit(1)
	}
	bmp, dib, err := bitmap.ReadBMP(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err.Error())
		os.Exit(1)
	}
	utils.PrintHeaderInfo(bmp, dib)
}

func parseHeaderArgs(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--help") || strings.HasPrefix(arg, "-h") {
			utils.PrintHeaderUsage()
		}
	}
}
