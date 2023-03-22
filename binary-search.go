package main

import "fmt"

func search() {

	var value int
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	down := 0
	top := len(list) - 1

	fmt.Println("----> Example of program with binary search <----")
	fmt.Println("Choose a number between 1 and 10: ")
	fmt.Scan(&value) //Receive the value from user input to var value
	//fmt.Println("You choose this value: ", value)
	var middle, guess int
	for down <= top {
		middle = (down + top) / 2
		guess = list[middle]

		if guess == value {
			// println(middle)
		}
		if guess > value {
			top = middle - 1
		} else {
			down = middle + 1
		}
	}
	fmt.Printf("Encontrado na posição %d de %d\n", middle, len(list))
}

func main() {
	search()
}
