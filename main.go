package main

import (
	"bitmap/internal/handlers"
	"bitmap/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		utils.DisplayUsage()
	}

	switch args[0] {
	case "header":
		// fmt.Println(strings.Split(args[1], ".bmp"))
		// fmt.Println("Length:", len(strings.Split(args[1], ".bmp")[0]))
		handlers.HandleHeader(args[1:])
	case "apply":
		handlers.HandleApply(args[1:])
	default:
		utils.DisplayUsage()
	}

}
