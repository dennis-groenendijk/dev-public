package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    needsLicense := false
    if kind == "car" || kind == "truck" {
        needsLicense = true
    }
	return needsLicense
    // no ternary conditional operator? (could be cleaner..)
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    vehicle := option2
    if option1 < option2 {
        vehicle = option1
    }
	return vehicle + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age < 3.0 {
		newPrice := originalPrice / 100 * 80
        return float64(newPrice)
	} else if age >= 3.0 && age < 10.0 {
		newPrice := originalPrice / 100 * 70
        return float64(newPrice)
	} else {
		newPrice := originalPrice / 2
        return float64(newPrice)
	}
}
