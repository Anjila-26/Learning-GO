package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	// rand.Intn(20) returns [0, 19]. Adding 1 shifts the range to [1, 20].
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// rand.Float64() returns [0.0, 1.0). Multiplying by 12.0 scales the range to [0.0, 12.0).
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	x := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    
	// rand.Shuffle shuffles the slice 'x' in place.
	rand.Shuffle(len(x), func(i int, j int) {
		x[i], x[j] = x[j], x[i]
	})
    
	// Return the shuffled slice.
	return x
}