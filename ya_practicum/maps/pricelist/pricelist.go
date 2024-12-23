package main

import "fmt"

func main() {
	// first exercise
	pricelist := map[string]int{
		"bread":      50,
		"milk":       100,
		"butter":     200,
		"souses":     500,
		"salt":       20,
		"cucumbers":  200,
		"chees":      600,
		"pork":       700,
		"extra pork": 900,
		"tomatoes":   250,
		"fish":       300,
		"hamon":      1500,
	}

	fmt.Println("1. Foodstuff over 500:")
	for k, v := range pricelist {
		if v > 500 {
			fmt.Printf("%v\n", k)
		}
	}

	// second exercise
	fmt.Println()
	order := []string{"bread", "extra pork", "chees", "cucumbers"}
	amount := 0
	fmt.Println("2. Order amount:")
	for o := range order {
		fmt.Println(pricelist[order[o]])
		amount += pricelist[order[o]]
	}

	fmt.Printf("Amount: %d", amount)
}
