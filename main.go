package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("Digite un número: ")
	fmt.Scan(&i)
	result, err := decToBin(i)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
