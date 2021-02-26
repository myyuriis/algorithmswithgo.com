package module01

import (
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	reversed := ""
	splitted := strings.Split(word, "")

	for i := len(splitted) - 1; i >= 0; i-- {
		reversed = reversed + splitted[i]
	}

	return reversed
}
