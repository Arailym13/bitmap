package utils

import (
	"fmt"
	"os"
)

func DisplayUsage() {
	fmt.Println("Usage: program [header|apply] [options]")
	os.Exit(1)
}
