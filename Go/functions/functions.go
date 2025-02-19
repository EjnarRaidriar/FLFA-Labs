package functions

import (
	"fmt"
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
