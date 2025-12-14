package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
    for i := 0; i < len(birdsPerDay); i++{
        count += birdsPerDay[i]
    }

    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// 1. Calculate the start index for the given week.
	// Since week numbers are 1-based (week 1 starts at index 0), 
    // we use (week - 1) * 7.
	start := (week - 1) * 7
    
	// 2. Calculate the end index for the given week (exclusive).
	// This is simply the week number multiplied by 7.
	end := week * 7
    
	// Check if the calculated indices are within the bounds of the slice.
    // The problem statement implies weeks are tracked completely, but this 
    // is a good practice check.
	if start < 0 || end > len(birdsPerDay) {
        // If the week is out of bounds, we assume the count is 0,
        // or you could panic/return an error depending on requirements.
		return 0 
	}
    
	// 3. Extract the sub-slice for the target week.
	weekCounts := birdsPerDay[start:end]
    
	// 4. Sum the counts in the sub-slice.
	totalBirds := 0
	for _, count := range weekCounts {
		totalBirds += count
	}

	return totalBirds
}

func FixBirdCountLog(birdsPerDay []int) []int {
	// The problem states the bird was present on the first day (index 0) 
	// and every second day thereafter (indices 0, 2, 4, 6, ...).
    
	// Iterate through the slice using an index i.
	for i := 0; i < len(birdsPerDay); i++ {
		// Check if the current day index 'i' is even (including day 0).
		if i%2 == 0 {
			// If the index is even, the missing bird was present.
			// Increment the count for that day by 1.
			birdsPerDay[i]++
		}
	}

	// Return the modified slice.
	return birdsPerDay
}
