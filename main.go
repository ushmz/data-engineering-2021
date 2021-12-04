package main

import (
	"fmt"
	"os"
)

var S = []interface{}{
	"a", "b", "a", "c", "b", "d", "e", "a", "g", "f", "b", "f", "f", "a", "c", "f", "a", "f",
}

func main() {
	// -------------------------------
	// Distinct elements
	// -------------------------------
	// Count elements
	file, err := os.Open("./content/benchmark.txt")
	if err != nil {
		panic(err)
	}
	count := DistinctElemCount(file)
	fmt.Println("Distinct elements :", count)
	file.Close()

	// Estimate elements
	f, err := os.Open("./content/benchmark.txt")
	if err != nil {
		panic(err)
	}
	est := DistinctElem(f)
	fmt.Println("Distinct elements estimation :", est)

	// -------------------------------
	// Frequent items
	// -------------------------------
	// Majority algorithm
	fmt.Println("Majority Algorythm")
	m, err := MajorityCount(S)
	if err != nil {
		fmt.Println("Majority:", err)
	} else {
		fmt.Println("Majority:", m)
	}

	// Lossy algorithm
	fmt.Println("Lossy Algorythm")
	m, err = LossyCounting(S, 4)
	if err != nil {
		fmt.Println("Majority:", err)
	} else {
		fmt.Println("Majority:", m)
	}

	// Space algorithm
	fmt.Println("SpaceSaving Algorythm")
	m, err = SpaceSaving(S, 4)
	if err != nil {
		fmt.Println("Majority:", err)
	} else {
		fmt.Println("Majority:", m)
	}
}
