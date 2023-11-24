package main

import (
	"fmt"
	"maps"
)

func main() {
	exampleMap := make(map[string]int)
	exampleMap["key"] = 10
	exampleMap["key2"] = 20
	exampleMap["key3"] = 30

	for key, value := range exampleMap {
		println(key, value)
	}

	fmt.Println("delete")
	maps.DeleteFunc(exampleMap, func(key string, value int) bool {
		return value > 20
	})

	for key, value := range exampleMap {
		println(key, value)
	}

	fmt.Println("copy")
	var example2Map = make(map[string]int)
	example2Map["key"] = 11
	maps.Copy(example2Map, exampleMap)

	for key, value := range example2Map {
		println(key, value)
	}

	e := maps.Equal(exampleMap, example2Map)
	fmt.Println("equal", e)
}
