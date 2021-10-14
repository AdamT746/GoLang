package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20,35,30}
	var ages = [3]int{20,35,30}

	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	names [1] = "Dave"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (similar to c# arrays)
	var scores = []int{10,20,50}
	scores [2] = 40
	scores = append(scores, 95)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Toad")
	fmt.Println(rangeOne)
}