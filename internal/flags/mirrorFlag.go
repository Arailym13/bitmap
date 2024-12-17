package flags

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func ParseMirrorFlag(flagValue string) []string {
	validOptions := map[string]string{
		"horizontal": "horizontal", "h": "horizontal", "horizontally": "horizontal", "hor": "horizontal",
		"vertical": "vertical", "v": "vertical", "vertically": "vertical", "ver": "vertical",
	}

	options := strings.Fields(flagValue)
	var parsedOptions []string
	for _, option := range options {
		normalizedOption := strings.ToLower(strings.TrimSpace(option))
		if validOption, exists := validOptions[normalizedOption]; exists {
			parsedOptions = append(parsedOptions, validOption)
		} else {
			fmt.Println("Error: Invalid mirror option")
			os.Exit(1)
		}
	}
	return parsedOptions
}

func MirrorHorizontal(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	mirroredImg := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mirroredImg.Set(width-1-x, y, img.At(x, y))
		}
	}
	return mirroredImg
}
