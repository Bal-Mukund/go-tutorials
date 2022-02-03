package main

import "fmt"

func main() {
	number := 5

	pointernum := &number
	*pointernum = *pointernum + 3
	fmt.Println(number)
}
