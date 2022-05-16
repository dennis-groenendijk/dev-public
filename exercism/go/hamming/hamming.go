package hamming

import "errors"

func Distance(a, b string) (int, error) {
    hd := 0
    
	if len(a) != len(b) {
        return -1, errors.New("The strands differ in length")
    }

    for i := range a {
        if a[i] != b[i] {
            hd++
        }
    }
	return hd, nil
}
