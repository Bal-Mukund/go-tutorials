package main

import "fmt"

func main() {
	var myarray [3]string
	myarray[0] = "mynumber"
	myarray[2] = "yournumber"
	fmt.Println(myarray)

	yourarray := [4]string{"your", "my", "our", "mine"}
	fmt.Println(yourarray)

}
