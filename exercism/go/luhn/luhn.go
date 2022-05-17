package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the provided input is a valid luhn string or not
func Valid(id string) bool {
	// 1. clean input: discard if len(id) =< 1 or non-digit characters and trim spaces
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	// 2.a convert string to a map
	digits := make([]int, len(id))

	for i, digit := range id {
		if digit < '0' || digit > '9' {
			return false
		}
		number, _ := strconv.Atoi(string(digit)) // ignoring the error
		digits[i] = number                       // swap the runes for integers
	}

	// 2.b start at last element and for the length of digits, count back (go left)
	// todo: doesn't calculate! see playground notes
	// for i := digits[len(digits)- 1]; i < len(digits); i -= 2 {
	// 	// 3.a double every second digit, starting from the right
	//     doubled := digits[i] * 2
	//     // 3.b doubling the number result in gt 9, subtract 9
	//         if doubled > 9 {
	//             doubled -= 9
	//         }
	// 		digits[i] = doubled
	// }

	for k, v := range digits {
		if k%2 != 0 {
			doubled := v * 2
			if doubled > 9 {
				doubled -= 9
			}
			v = doubled
		}
	}

	// 4. sum all digits
	sumOfDigits := 0
	for _, digit := range digits {
		sumOfDigits += digit
	}

	// 5. if the sum is divisible by 10, return true
	return sumOfDigits%10 == 0
}
