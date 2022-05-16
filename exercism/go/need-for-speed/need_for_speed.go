package speed

// TODO: define the 'Car' type struct
type Car struct {
    speed 			int
    batteryDrain 	int
    battery 		int
    distance 		int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car {
        speed: 			speed,
        batteryDrain: 	batteryDrain,
        battery:		100,
    }
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
    return Track {
        distance: distance,
    }
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
    newDistance := car.distance + car.speed
    newBattery := car.battery - car.batteryDrain

    if newBattery > 0 {
        car.battery = newBattery
        car.distance = newDistance
    }
	return car
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    maxDistance := car.battery / car.batteryDrain * car.speed
    return maxDistance >= track.distance
	panic("Please implement the CanFinish function")
}
