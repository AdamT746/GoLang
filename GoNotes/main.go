package main

import (
	"fmt"
	"sort"
	
)

func main () {
	//greeting := "hello there"

	//fmt.Println(strings.Contains(greeting, "hello")) //true false if string contains
	//fmt.Println(strings.ReplaceAll( greeting,"hello", "alright")) // replace parts of strings
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "ll"))
	//fmt.Println(strings.Split(greeting, " "))

	//This doesnt change the string but returns a new one
	//fmt.Println("Original value = ", greeting)

	ages := []int{20,35,40,22,43,50,69,31}
	sort.Ints(ages)
	fmt.Println(ages)
}