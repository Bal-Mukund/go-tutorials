package main

import (
	"fmt"
	"sort"
)

func main() {
	food := []string{}

	food = append(food, "meat")
	food[0] = "veg"
	fmt.Println(food)

	scores := make([]int, 4)
	scores[0] = 92
	scores[1] = 45
	scores[2] = 56
	scores[3] = 78

	scores = append(scores, 675, 343, 2453) // realocate the memory while using slices
	sliced_score := append(scores[:3])

	sliced_score[0] = 0 // pointers

	fmt.Println(sliced_score)
	sort.Ints(scores)
	fmt.Println(scores)

	// remove a value from a slice using an index

	index := 1

	role := []string{"actor", "director", "music", "dance"}
	role = append(role[:index], role[index+1:]...)
	fmt.Println(role)

}
