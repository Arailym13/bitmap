package utils

import (
	"fmt"
	"os"
)

func PrintApplyUsage() {
	fmt.Println(`
Usage:
  bitmap apply [options] <source_file> <output_file>

The options are:
  -h, --help      prints program usage information
  --mirror        mirrors a bitmap image either horizontally or vertically
                   available options:
                      horizontal, h, horizontally, hor
                      vertical, v, vertically, ver

  --filter        apply various filters to images
                    available filters:
                      blue, red, green, grayscale, negative, pixelatesize, blur

  --rotate        rotates a bitmap image by a specified angle
                    available rotations:
                    right (clockwise), 90, 180, 270, left, -90, -180, -270

  --crop          trims a bitmap image according to specified parameters
                    "--crop=OffsetX-OffsetY-Width-Height"
                  You can specify crop without width and height
  ...`)
	os.Exit(1)
}
