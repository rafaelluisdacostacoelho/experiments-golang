package main

import (
	"fmt"
)

func main() {
	items := []int{1, 2, 3, 4}

	items = append(items, 5)
	items = append(items, 6)

	// cap() exibe a capacidade do array/slice
	fmt.Println("Capacity:", cap(items))

	// len() exibe a quantidade de elementos no array
	fmt.Println("Length:", len(items))
}
