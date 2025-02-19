package functions

import (
	"fmt"
	"unicode"
)

func PrintRuneList(list []rune) {
	fmt.Print("[")
	l := len(list)
	for i, e := range list {
		if i+1 == l {
			fmt.Printf("%c", e)
			break
		}
		fmt.Printf("%c ", e)
	}
	fmt.Print("]")
}

func RuneSliceToStringSlice(runes []rune) []string {
	var stringSlice []string
	for _, r := range runes {
		stringSlice = append(stringSlice, string(r))
	}
	return stringSlice
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func HasMultipleUpper(s string) bool {
	upperCount := 0
	for _, r := range s {
		if unicode.IsUpper(r) {
			upperCount++
		}
		if upperCount > 1 {
			return true
		}
	}
	return false
}
