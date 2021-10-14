package main

import "fmt"

func main() {

	//Strings
	var NameOne string = "Man"
	var NameTwo = "Woman"
	var NameThree string
	
	NameThree = "Child"

	NameFour := "Auntie"

	fmt.Println(NameOne, NameTwo, NameThree, NameFour)


	//Integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne,ageTwo,ageThree)

	//Bits & Mem
	//var numOne int8 = 25  
	//var numTwo int8 = -128
	//var numThree uint = 25 //Doesn't allow negative numbers

	var scoreOne float32 = -1.5
	var scoreTwo float64 = 7645634634534.3
	scoreThree := 5.3 // infers the best type to use (float64)

}
