package main

import (
	"fmt"
	"github.com/quagmt/udecimal"
)

func main() {
	// Create a new decimal number
	a, _ := udecimal.NewFromInt64(123456, 3)              // a = 123.456
	b, _ := udecimal.NewFromInt64(-123456, 4)             // b = -12.3456
	c, _ := udecimal.NewFromFloat64(1.2345)               // c = 1.2345
	d, _ := udecimal.Parse("4123547.1234567890123456789") // d = 4123547.1234567890123456789

	// Basic arithmetic operations
	fmt.Println(a.Add(b)) // 123.456 - 12.3456 = 111.1104
	fmt.Println(a.Sub(b)) // 123.456 + 12.3456 = 135.8016
	fmt.Println(a.Mul(b)) // 123.456 * -12.3456 = -1524.1383936
	fmt.Println(a.Div(b)) // 123.456 / -12.3456 = -10
	fmt.Println(a.Div(d)) // 123.456 / 4123547.1234567890123456789 = 0.0000299392722585176

	// Rounding
	fmt.Println(c.RoundBank(3)) // banker's rounding: 1.2345 -> 1.234
	fmt.Println(c.RoundHAZ(3))  // half away from zero: 1.2345 -> 1.235
	fmt.Println(c.RoundHTZ(3))  // half towards zero: 1.2345 -> 1.234
	fmt.Println(c.Trunc(2))     // truncate: 1.2345 -> 1.23
	fmt.Println(c.Floor())      // floor: 1.2345 -> 1
	fmt.Println(c.Ceil())       // ceil: 1.2345 -> 2

	// Display
	fmt.Println(a.String())         // 123.456
	fmt.Println(a.StringFixed(10))  // 123.4560000000
	fmt.Println(a.InexactFloat64()) // 123.456
}
