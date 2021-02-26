package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// Using strings.Split function (reverse for loop)
//func Reverse(word string) string {
//	reversed := ""
//	splitted := strings.Split(word, "")
//
//	for i := len(splitted) - 1; i >= 0; i-- {
//		reversed = reversed + splitted[i]
//	}
//	return reversed
//}

// Using string() function to convert byte to string (reverse for loop)
//func Reverse(word string) string {
//	reversed := ""
//
//	for i := len(word)-1; i >= 0; i-- {
//		reversed = reversed + string(word[i])
//	}
//	return reversed
//}

// Using string builder
//func Reverse(word string) string {
//	var sb strings.Builder
//
//	for i := len(word)-1; i >= 0; i-- {
//		sb.WriteByte(word[i])
//	}
//	return sb.String()
//}

// Using normal order (not reversed for loop)
//func Reverse(word string) string {
//	rev := ""
//
//	for i := 0; i < len(word); i++ {
//		rev = string(word[i]) + rev
//	}
//	return rev
//}

// Using for range to account for multi-byte characters (test: 日本語)
func Reverse(word string) string {
	rev := ""

	for _, v := range word {
		rev = string(v) + rev
	}
	return rev
}