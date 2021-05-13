// Created by: haokai
// Created on: May 2021
// This program displays, "Volume-of-a-sphere"

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function displays currency
	accountingFormater := accounting.Accounting{Precision: 2}
	// This function does addition
	var radius float64
	var volume float64
	// input
	fmt.Println("This program is about Volume of a sphere program.")
	fmt.Println()
	fmt.Print("Enter the number of radius(cm): ")
	fmt.Scanln(&radius)
	fmt.Println()
	// process
	volume = radius * radius * radius * 3.1415926535897932384626433 * 4 / 3
	// output
	fmt.Println("Volume≈ ", accountingFormater.FormatMoney(volume), "cm³")
	fmt.Println("\nDone.")
}
