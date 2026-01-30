package main

import (
	"fmt"
	"maps"
)

func compareMaps() {
	my_map_a := map[string]int{
		"hello": 5,
		"world": 10,
	}
	my_map_b := map[string]int{
		"world": 10,
		"hello": 50,
	}

	fmt.Println(maps.Equal(my_map_a, my_map_b))

}
