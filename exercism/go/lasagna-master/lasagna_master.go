package lasagna

// PreparationTime calculates the time it takes to prepare lasagna by multiplying
// the number of layers with the average preparation time per layer in minutes
func PreparationTime(layers []string, avgTime int) int {
    if avgTime == 0 {
        avgTime = 2
    }
    prepTime := len(layers) * avgTime
    return prepTime
}

// Quantities determines the amount of noodles and sauce needed to make the lasagna
func Quantities(layers []string) (noodles int, sauce float64) {
    grams := 0
    liters := 0.0

    for _, v := range layers {
        if v == "noodles" {
            grams++
        } else if v == "sauce" {
        	liters++
        }
    }

    noodles = grams * 50
    sauce = liters * 0.2

    return noodles, sauce
}

// AddSecretIngredient compares two lists of ingredients and copies the missing ingredient
// from one to the other
func AddSecretIngredient(friendslist, myList []string) {
    for i, v := range friendslist {
        if i == len(friendslist) - 1 {
            myList = append(myList[:len(myList) - 1], v)
        }
    }
}

// ScaleRecipe calculates the amounts of ingredients needed for more than 2 portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
    var scaledQuantities []float64
    
    for i, v := range quantities {
        if len(scaledQuantities) < len(quantities) {
          	scaledValue := (v / 2.0) * float64(portions)
        	scaledQuantities = append(scaledQuantities, scaledValue)
        	i++  
        }
    }
	return scaledQuantities
}
