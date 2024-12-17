package main

import (
	"bitmap/pkg"
	"fmt"
	"os"
)

func main() {
	bmpFile := "sample.bmp"
	fmt.Println(bmpFile)
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "header" {
		pkg.PrintHeaderInfo(bmpFile)
	}
}
