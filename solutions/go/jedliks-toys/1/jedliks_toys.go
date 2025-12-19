package jedlik

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
// CanFinish checks if the car can finish the track without running out of battery
func (c Car) CanFinish(trackDistance int) bool {
	// Number of drives needed (ceiling division)
	drivesNeeded := trackDistance / c.speed
	if trackDistance%c.speed != 0 {
		drivesNeeded++
	}

	requiredBattery := drivesNeeded * c.batteryDrain
	return requiredBattery <= c.battery
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
