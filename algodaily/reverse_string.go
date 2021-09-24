package algodaily

import "strings"

//ReverseString takes a string and returns the string reversed.
func ReverseString(phrase string) string {
	start := 0
	end := len(phrase) - 1
	splitted := strings.Split(phrase, "")

	for start <= end {
		temp := splitted[start]
		splitted[start] = splitted[end]
		splitted[end] = temp

		start++
		end--
	}

	return strings.Join(splitted, "")
}
