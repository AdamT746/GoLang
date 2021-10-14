package main
import "fmt"

func main () {

	age := 30
	name := "dave"

	//print + print line
	fmt.Print("hello, ") //doesnt make a new line
	fmt.Print("World \n") // \n creates a new line
	fmt.Print("New")

	fmt.Println() // auto creates a new line

	fmt.Println("My age is", age , "and my name is", name ) // printing variables


	//formatted string 
	fmt.Printf("My age is %v and my name is %v \n", age, name) //%_ = format specifier
	fmt.Printf("age is of type %T \n", age) //%T finds the type of the variable
	fmt.Printf("you scored %0.1f points! \n", 23.456) //%f is full number 0.1f is one decimal place

	//Sprintf (save formatted strings)
	var string = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", string)
}