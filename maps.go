package main

import "fmt"

// Map
func maps() {
	// String map
	menu := map[string]float64{
		"samsung": 12.66,
		"iphone":  15.66,
		"Oppo":    10.32,
		"vivo":    11.11,
	}

	fmt.Println(menu)
	fmt.Println(menu["vivo"])

	// Looping Map
	for key, value := range menu {
		fmt.Println(key, "- ", value)
	}

	// Integer Map
	phoneNumber := map[int]string{
		792826217: "Nyan",
		792826218: "Lin",
		792826219: "Htoo",
	}

	fmt.Println(phoneNumber)
	fmt.Println(phoneNumber[792826217])

	// Changing Value in Map
	phoneNumber[792826218] = "LinLin"
	fmt.Println(phoneNumber)

	phoneNumber[792826219] = "HtooHtoo"
	fmt.Println(phoneNumber)
}
