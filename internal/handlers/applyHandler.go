package handlers

import (
	"bitmap/internal/bitmap"
	"bitmap/utils"
	"fmt"
	"os"
	"strings"
)

type ApplyOption struct {
	Name  string
	Value string
}

func HandleApply(args []string) {
	fmt.Println("-----Apply-----")
	options, inputFile, _, err := parseApplyArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	_, _, err = bitmap.ReadBMP(inputFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, option := range options {
		switch option.Name {
		case "mirror":
			fmt.Println("Mirror with value", option.Value)
		default:
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
	}

}

func parseApplyArgs(args []string) ([]ApplyOption, string, string, error) {
	options := []ApplyOption{}
	var inputFile, outputFile string
	var fileArgs []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--mirror") || strings.HasPrefix(arg, "-mirror") {
			var flagValue string
			if arg[1] == '-' {
				flagValue = strings.TrimPrefix(arg, "--mirror=")
			} else {
				flagValue = strings.TrimPrefix(arg, "-mirror=")
			}
			options = append(options, ApplyOption{"mirror", flagValue})
		} else if !strings.HasPrefix(arg, "-") {
			fileArgs = append(fileArgs, arg)
		} else {
			if len(fileArgs) != 2 {
				return nil, "", "", fmt.Errorf("'apply' requires exactly two positional arguments: input file and output file")
			}
			return nil, "", "", fmt.Errorf("unexpected argument: %s", arg)
		}
	}
	inputFile = fileArgs[0]
	outputFile = fileArgs[1]

	if !utils.IsValidBMPFile(inputFile) || !utils.IsValidBMPFile(outputFile) {
		return nil, "", "", fmt.Errorf("one or both of these files are not in bmp format")
	}
	return options, inputFile, outputFile, nil
}
