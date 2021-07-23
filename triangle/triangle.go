// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if a+b >= c && a+c >= b && c+b >= a && a*b*c > 0 && a*b*c < math.Inf(1) {
		if a == b && b == c {
			return Equ
		} else if a == b || b == c || c == a {
			return Iso
		}
		return Sca
	}
	return NaT
}
