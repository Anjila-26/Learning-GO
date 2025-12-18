package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
    if minutes == 0 {
        return len(layers) * 2
    } else {
        return len(layers) * minutes
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
	for i := 0; i < len(layers); i++{
        switch{
            case layers[i] == "noodles":
            	noodles += 50
            case layers[i] == "sauce":
            	sauce += 0.2
        }
    }

    return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
    // 1. Get the secret ingredient from the last element of friendsList.
    // Index of the last element is len(friendsList) - 1.
    secretIngredient := friendsList[len(friendsList)-1]
    
    // 2. Get the index of the last element of myList (which is the "?").
    myListLastIndex := len(myList) - 1
    
    // 3. Replace the last element of myList with the secret ingredient.
    myList[myListLastIndex] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	// The original recipe is for 2 portions.
	const originalPortions = 2

	// 1. Calculate the scaling factor. 
	// Convert 'portions' to float64 for correct division.
	scalingFactor := float64(portions) / originalPortions

	// 2. Create a new slice to hold the scaled amounts. 
	// This prevents modifying the original 'quantities' slice.
	scaledQuantities := make([]float64, len(quantities))

	// 3. Iterate through the original quantities, scale each one, and copy to the new slice.
	for i, amount := range quantities {
		// Calculate the new amount.
		scaledAmount := amount * scalingFactor
        
		// Store the result in the new slice.
		scaledQuantities[i] = scaledAmount
	}

	return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
