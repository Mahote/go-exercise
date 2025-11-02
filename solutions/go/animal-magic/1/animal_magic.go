package chance

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return randInt(1,20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return randFloat(0.0,12.0)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animalSlice := []string{"Lion", "Tiger", "Bear", "Wolf", "Fox", "Eagle", "Hawk", "Falcon"}
	return shuffleSlice(animalSlice)
}
