package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(opt)
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
