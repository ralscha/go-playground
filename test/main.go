package main

import "fmt"

func DiscountedPrice(price, discountPercent float64) (float64, error) {
	switch {
	case discountPercent == 0:
		return price, nil

	case discountPercent < 0:
		return 0, fmt.Errorf(
			"invalid negative discount percentage: %.2f",
			discountPercent)

	case discountPercent > 100:
		return 0, fmt.Errorf(
			"invalid discount percentage greater than 100: %.2f",
			discountPercent)

	default:
		discount := price * (discountPercent / 100)
		return price - discount, nil
	}
}
