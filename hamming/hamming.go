package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	// 0th condition to check equal length
	if len(a) != len(b) {
		return 0, fmt.Errorf("disallowed: length of both strings should equal")
	}

	var ham int = 0
	// if both strings are identical
	if a == b {
		return 0, nil
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				ham += 1
			} else {
				continue
			}
		}
		return ham, nil
	}

}
