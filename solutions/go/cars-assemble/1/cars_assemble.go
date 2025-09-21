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
func CalculateCost(carsCount int) uint {
	//cost of production for one car
	var productionCostPerCar uint = 10000
	//cost of production for 10 cars
	var productionCostPerTenCar uint = 95000
	var cost uint = 0

	if carsCount < 10 {
		cost = uint(carsCount) * 10000
	} else {
		cost = uint(carsCount/10)*productionCostPerTenCar + uint(carsCount%10)*productionCostPerCar
	}
	return cost
}
