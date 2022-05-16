package romannumerals

import (
	"errors"
    "strconv"
)

// ToRomanNumeral converts the provided Arabic number to a Roman numeral.
// If the provided input is egt 3000, 0 or a negative number it results in an error msg.
func ToRomanNumeral(input int) (string, error) {
	maxNumber := 3000
    if input > maxNumber || input <= 0 {
        return strconv.Itoa(input), errors.New(": can't convert this number")
    }

    conversions := []struct {
        number int
        numeral string
    }{
    	{1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    output := ""
    for _, conversion := range conversions {
        for input >= conversion.number {
            output += conversion.numeral
            input -= conversion.number
        }
    }
	return output, nil
}
