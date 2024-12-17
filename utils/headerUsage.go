package utils

import (
	"fmt"
	"os"
)

func PrintHeaderUsage() {
	fmt.Println(`
Usage:
  bitmap header <source_file>

Description:
  Prints bitmap file header information`)
	os.Exit(1)
}
