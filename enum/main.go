package main

import (
	"enum/vehicle"
	"errors"
	"fmt"
	"log"
)

/*
type FakeVehicle {
	vehicle.Type
}
*/

func main() {

	v := vehicle.Values.Truck
	rate, err := calculateInsuranceRate(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rate)

	// rate, err := calculateInsuranceRate(&FakeVehicle{})
}

func calculateInsuranceRate(v vehicle.Type) (float64, error) {
	switch v {
	case vehicle.Values.Motorcycle:
		return 0.05, nil
	case vehicle.Values.Car:
		return 0.2, nil
	case vehicle.Values.Bus:
		return 0.3, nil
	case vehicle.Values.Truck:
		// We can even invoke some methods on the concrete value
		return 0.3 * vehicle.Values.Truck.FetchSomeData(), nil
	default:
		return 0, errors.New("vehicle type undefined")
	}
}
