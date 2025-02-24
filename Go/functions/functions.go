package functions

import (
	"cmp"
	"fmt"
	"slices"
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

func KeyList[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func RemoveDuplicates[T comparable](list []T) []T {
	keys := make(map[T]bool)
	newList := []T{}
	for _, value := range list {
		if _, ok := keys[value]; !ok {
			keys[value] = true
			newList = append(newList, value)
		}
	}
	return newList
}

func ContainsMap[K cmp.Ordered, V any](list []map[K]V, m map[K]V) bool {
	for _, element := range list {
		if slices.Compare(KeyList(element), KeyList(m)) == 0 {
			return true
		}
	}
	return false
}
