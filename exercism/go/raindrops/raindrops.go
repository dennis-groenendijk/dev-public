package raindrops

import "strconv"

func Convert(number int) (str string) {
    if number % 3 == 0 {
        str += "Pling"
    }
	if number % 5 == 0 {
    	str += "Plang"
    }
	if number % 7 == 0 {
    	str += "Plong"
    }
	if str == "" {
    	str = strconv.Itoa(number)
    }
	return
}
