package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper Function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func promptOpt(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - Add item, s - save bill, t - add tip): ", reader)
	// fmt.Println(opt)

	// Switch Statements
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Enter price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number")
			promptOpt(b)
		}

		b.addItems(name, p)

		fmt.Println("Added Item: ", name, price)
		promptOpt(b)
	case "t":
		tip, _ := getInput("Enter tip ammount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number")
			promptOpt(b)
		}
		b.updateTip(t)
		fmt.Println("Added Tip: ", tip)
		promptOpt(b)
	case "s":
		b.save()
		fmt.Println("You save bill - ", b.name)
	default:
		fmt.Println("That Choice is not a valid!")
		promptOpt(b)
	}
}

// Save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to the file")
}

func createBill() bill {
	// bufio     =>  to read user input  info
	// os.Stdin  => Standard Input from terminal
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a name of bill: ")
	// name, _ := reader.ReadString('\n') // \n is coming word after read
	// name = strings.TrimSpace(name)     // Trim space

	name, _ := getInput("Create a name of bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}
