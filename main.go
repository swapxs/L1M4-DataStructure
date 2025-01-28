package main

import "fmt"

func sol(cap1, cap2, t int) {
	var bigJug, smallJug, steps int

	for bigJug != t && smallJug != t {

		if bigJug == 0 {
			bigJug = cap1
			fmt.Printf("Step %v: Fill Big Jug to %v gallons\n", steps, bigJug)
		} else if smallJug != cap2 {
			emptySpace := cap2 - smallJug
			if bigJug <= emptySpace {
				smallJug += bigJug
				bigJug = 0
			} else {
				bigJug -= emptySpace
				smallJug = cap2
			}
			fmt.Printf("Step %v: Pour Big Jug into Small Jug,\n\t Big Jug: %v, Small Jug: %v\n", steps, bigJug, smallJug)
		}

		if smallJug == cap2 {
			smallJug = 0
			fmt.Printf("Step %v: Empty Small Jug\n", steps)
		} else if bigJug == 0 {
			bigJug = cap1
			fmt.Printf("Step %v: Fill Big Jug to %v gallons\n", steps, bigJug)
		}

		steps++
	}

	if bigJug == t {
		fmt.Printf("Big Jug has %v gallons of water\n", bigJug)
	} else if smallJug == t {
		fmt.Printf("Small Jug has %v gallons of water\n", smallJug)
	}
}

func main() {
	bigJugCap := 4
	smallJugCap := 3
	t := 2

	fmt.Printf("=========== The Water Jug Problem – Count Min Steps ===========\n")
	fmt.Printf("Size of Big Jug: %v\nSize of Small Jug: %v\nTarger: %v\n\n", bigJugCap, smallJugCap, t)
	sol(bigJugCap, smallJugCap, t)

}
