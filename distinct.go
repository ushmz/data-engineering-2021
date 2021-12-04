package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/twmb/murmur3"
)

func DistinctElemCount(file *os.File) int {
	ipmap := map[string]bool{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		// Clean dataset (Not so important)
		nodes := strings.Split(line, " ")
		ipstr := strings.Join(nodes[len(nodes)-5:len(nodes)-1], ".")

		if _, ok := ipmap[ipstr]; !ok {
			ipmap[ipstr] = true
		}
	}
	return len(ipmap)
}

func DistinctElem(file *os.File) int {
	var min uint64
	var max uint64
	hasher := murmur3.SeedNew64(uint64(42))

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		// Clean dataset (Not so important)
		nodes := strings.Split(line, " ")
		ipstr := strings.Join(nodes[len(nodes)-5:len(nodes)-1], ".")

		hasher.Write([]byte(ipstr))
		sum := hasher.Sum64()
		if min == 0 {
			min = sum
		}
		if max == 0 {
			max = sum
		}
		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max, min)
	// m := math.Pow(256, 4)
	// fmt.Println("Float64", m)
	// fmt.Println("uiint64", uint64(m))
	s := ((max) / min) - uint64(1)
	fmt.Println("uint64", s)
	return int(s)
}
