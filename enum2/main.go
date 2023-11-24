package main

import (
	"fmt"
	"github.com/orsinium-labs/enum"
)

type Color enum.Member[string]

var (
	Red    = Color{"red"}
	Green  = Color{"green"}
	Blue   = Color{"blue"}
	Colors = enum.New(Red, Green, Blue)
)

type Color2 enum.Member[string]

var (
	b       = enum.NewBuilder[string, Color2]()
	Red2    = b.Add(Color2{"red"})
	Green2  = b.Add(Color2{"green"})
	Blue2   = b.Add(Color2{"blue"})
	Colors2 = b.Enum()
)

func main() {
	parsed := Colors.Parse("red")
	fmt.Println(*parsed == Red)
	fmt.Println(*parsed == Green)

	for _, color := range Colors.Members() {
		fmt.Println(color)
		f(color)
	}

	for _, color2 := range Colors2.Members() {
		fmt.Println(color2)
		if color2 == Red2 {
			fmt.Println("red")
		} else if color2 == Green2 {
			fmt.Println("green")
		} else if color2 == Blue2 {
			fmt.Println("blue")
		}
	}
}

func f(color Color) {
	if !Colors.Contains(color) {
		panic("invalid color")
	}
	fmt.Println("ok")
}
