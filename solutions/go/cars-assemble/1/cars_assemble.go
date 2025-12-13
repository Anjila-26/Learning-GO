package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// 1. Integer division (/) finds the number of full groups of ten (the quotient).
	groupsOfTen := carsCount / 10 

	// 2. Modulo (%) finds the number of remaining individual cars (the remainder).
	individualCars := carsCount % 10 

	// Calculate the total cost and convert the final result to the required 'uint' type.
	totalCost := groupsOfTen * 95000 + individualCars * 10000

	return uint(totalCost)
}
