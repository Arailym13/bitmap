package utils

import "strings"

func IsValidBMPFile(filename string) bool {
	if !strings.HasSuffix(filename, ".bmp") {
		return false
	}
	if strings.HasSuffix(filename, ".bmp") && len(strings.Split(filename, ".bmp")[0]) == 0 {
		return false
	}
	return true
}
