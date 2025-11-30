package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) (res int) {
	res = len(layers) * avgPrepTime
	// Default average prep time is 2 minutes per layer
	if avgPrepTime == 0 {
		res = len(layers) * 2
	}
	return
}	
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}


// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, portions int) []float64 {
    if len(recipe) == 0 {
        return []float64{}
    }
    scaled := make([]float64, len(recipe))
    scaleFactor := float64(portions) / 2.0
    for i, ingredient := range recipe {
        scaled[i] = ingredient * scaleFactor
    }
    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
