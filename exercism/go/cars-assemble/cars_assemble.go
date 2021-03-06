package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pr := float64(productionRate)
	percentage := successRate / 100.00
	producedSuccessful := pr * percentage
	return producedSuccessful
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	pr := float64(productionRate)
	percentage := successRate / 100.00
	producedSuccessful := pr * percentage / 60.00
	producedSuccessfulPerMinute := int(producedSuccessful)
	return producedSuccessfulPerMinute
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groupsOfTenCost := carsCount / 10 * 95000
	remainingCarsCost := carsCount % 10 * 10000
	totalCost := uint(groupsOfTenCost + remainingCarsCost)
	return totalCost
}
