package speed

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
    return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain{
        // 1. Update distance: Increase distance by the car's speed.
		car.distance += car.speed

		// 2. Update battery: Reduce battery by the car's drain amount.
		car.battery -= car.batteryDrain
    }

    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// 1. Calculate how many times the car needs to "Drive" to cover the track distance.
	numberOfDrives := track.distance / car.speed

	// 2. Calculate the total battery needed for those drives.
	totalBatteryCost := numberOfDrives * car.batteryDrain

	// 3. If the total cost is less than or equal to the current battery, return true.
	return car.battery >= totalBatteryCost
}
