package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
    if index < 0 || index > len(slice)-1 {
        return 0, false
    }
	return slice[index], true
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < 0 || index > len(slice)-1 {
        slice = append(slice, value)
    } else {
    	slice[index] = value
    }
	return slice
	panic("Please implement the SetItem function")
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
    if length <= 0 {
        return nil
    }
	slice := make([]int, length)
    for i := range slice {
        slice[i] = value
    }
	return slice
	panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index < 0 || index > len(slice)-1 {
        return slice
    }
	/* The new slice is build up by concatenating the first half 
	(from first index up to the "index" input),
    followed by the second half (the first index after the "index"
	input, till the end of this slice)
    The "..." operator is what executes the concatenation. */
	return append(slice[:index], slice[index + 1:]...)
	panic("Please implement the RemoveItem function")
}
