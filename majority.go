package main

import (
	"errors"
	"fmt"
)

var S = []interface{}{"a", "b", "a", "c", "b", "d", "e", "a", "g", "f", "b", "f", "f", "a", "c", "f", "a", "f"}

func majorityCount(stream []interface{}) (interface{}, error) {
	var majority interface{}
	counter := 0

	for _, v := range stream {
		if majority == nil {
			majority = v
			counter++
		} else if majority == v {
			counter++
		} else {
			counter--
			if counter == 0 {
				majority = nil
			}
		}
		fmt.Println("Item:", majority, "\t Counter: ", counter)
	}

	if majority == nil {
		return nil, errors.New("There is no majority elements")
	}

	return majority, nil
}

func lossyCounting(stream []interface{}, k int) (interface{}, error) {
	delta := 0
	t := map[interface{}]int{}

	for i, v := range stream {
		if val, ok := t[v]; ok {
			t[v] = val + 1
		} else {
			t[v] = delta + 1
		}

		if (i+1)/k != delta {
			delta = (i + 1) / k
			for key, value := range t {
				if value <= delta {
					delete(t, key)
				}
			}
		}
		fmt.Println("Index:", i, "\t: ", t)
	}
	return t, nil
}

func main() {
	// Majority algorithm
	fmt.Println("Majority Algorythm")
	m, err := majorityCount(S)
	if err != nil {
		fmt.Println("Majority:", err)
	} else {
		fmt.Println("Majority:", m)
	}

	// Lossy algorithm
	fmt.Println("Lossy Algorythm")
	m, err = lossyCounting(S, 4)
	if err != nil {
		fmt.Println("Majority:", err)
	} else {
		fmt.Println("Majority:", m)
	}

	// Space algorithm
}
