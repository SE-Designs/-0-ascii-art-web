package app

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func First(input string, font string) (string, error) {
	banner, err := openFile(font)
	if err != nil {
		return "", errors.New(http.StatusText(http.StatusInternalServerError))
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")

	f := func(c rune) bool {
		return c == '\n'
	}

	banner_lines := strings.FieldsFunc(banner, f)
	input_lines := strings.Split(input, "\n")

	if len(input_lines) > len(input) {
		input_lines = input_lines[1:]
	}

	ascii_out := printOutput(input_lines, banner_lines)

	return ascii_out, nil
}

func openFile(font string) (string, error) {
	var banner []byte
	var err error
	// taking second argument as a font-style
	if font == "" {
		banner, err = os.ReadFile("app/standard.txt")
	} else if font == "shadow" {
		banner, err = os.ReadFile("app/shadow.txt")
	} else if font == "standard" {
		banner, err = os.ReadFile("app/standard.txt")
	} else if font == "thinkertoy" {
		banner, err = os.ReadFile("app/thinkertoy.txt")
	} else {
		fmt.Println("ERROR: cannot find the banner file")
		return "", errors.New("ascii-art-web/app/app/go: ERROR: cannot find the banner file")
	}
	if err != nil {
		fmt.Println("ERROR: cannot find the banner file")
		return "", errors.New("ascii-art-web/app/app/go: ERROR: cannot find the banner file")
	}
	// checking the correctness of the file
	hash_std := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	hash_thinkertoy := []byte{165, 123, 238, 196, 63, 222, 103, 81, 186, 29, 48, 73, 91, 9, 38, 88, 160, 100, 69, 47, 50, 30, 34, 29, 8, 195, 172, 52, 169, 220, 18, 148}
	hash_shadow := []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
	h := sha256.New()
	h.Write(banner)
	if string(h.Sum(nil)) != string(hash_std) && string(h.Sum(nil)) != string(hash_thinkertoy) && string(h.Sum(nil)) != string(hash_shadow) {
		fmt.Println("ERROR: corrupted banner file")
		return "", errors.New("ascii-art-web/app/app/go: ERROR: cannot find the banner file")
	}
	if len(banner) == 0 {
		fmt.Println("ERROR: empty banner")
		return "", errors.New("ascii-art-web/app/app/go: ERROR: cannot find the banner file")
	}
	return string(banner), nil
}

// check for invalid characters
func CheckString(s string) bool {
	for _, item := range s {
		if item == '\n' || item == '\r' {
			continue
		}
		if !(item >= ' ' && item <= '~') {
			fmt.Println("ERROR: invalid characters")
			return false
		}
	}
	return true
}

// printing the output
func printOutput(s []string, b []string) string {
	output := ""

	for _, word := range s {
		if word == "" {
			output += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, val := range word {
				if val >= ' ' && val <= '~' {
					output += b[8*(int(val)-32)+i]
					// fmt.Print(b[8*(int(val)-32)+i])
				}
			}
			output += "\n"
			// fmt.Print("\n")
		}
	}
	return output
}
