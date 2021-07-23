package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	// 0th condition to check equal length
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("disallowed: length of both strings should equal")
	}

	var ham int = 0
	var idx int = 0

	stringToRunes := []rune(b)

	for _, value := range a {
		if value != stringToRunes[idx] {
			ham += 1
		}
		idx++
	}
	return ham, nil
}
