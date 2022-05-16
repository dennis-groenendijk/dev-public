package triangle

import "math"

type Kind int

const (
    NaT Kind = iota	// not a triangle
    Equ 			// equilateral
    Iso 			// isosceles
    Sca 			// scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	var k Kind

    if math.IsNaN(a + b + c) || a + b <= c || a + c <= b || b + c <= a {
        k = NaT
    } else if a == b && a == c {
        k = Equ 
    } else if a == b || a == c || b == c {
        k = Iso
    } else {
    	k = Sca
    }
	return k
}
