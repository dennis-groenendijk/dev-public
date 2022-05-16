package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    rot := OvenTime - actualMinutesInOven
    return rot
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    timePerLayer := 2
    prep := numberOfLayers * timePerLayer
    return prep
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    timePerLayer := 2
	prep := numberOfLayers * timePerLayer
	totalTime := prep + actualMinutesInOven
	return totalTime
}

func main() {
    RemainingOvenTime(30)
	PreparationTime(2)
    ElapsedTime(3, 20)
}