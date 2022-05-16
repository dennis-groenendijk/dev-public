package twofer

// ShareWith parses the input "name" to the output phrase and checks if a name is provided or not
// if not, the name is replaced by "you".
func ShareWith(name string) string {
	if name == "" {
        return "One for you, one for me."
    } else {
    	return "One for " + name + ", one for me."
    }
}
