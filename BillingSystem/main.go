package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	Input, err := r.ReadString('\n') //Reader named r

	return strings.TrimSpace(Input), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := GetInput("Create a new bill name:", reader)

	b := newBill(name) //Bill named b
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := GetInput("Choose an option (a - add item, s - save bill, t - add tip):", reader)

	switch opt {
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price: ", reader)

		pricefloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, pricefloat)

		fmt.Println("Item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := GetInput("Enter tip amount: ", reader)

		tipfloat, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.UpdateTip(tipfloat)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You chose to save your bill", b.name)
	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}
func main() {
	mybill := createBill()
	promptOptions(mybill)

}
