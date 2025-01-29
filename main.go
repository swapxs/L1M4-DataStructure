// Implementating the program using BFS
package main

import "fmt"

type State struct {
	jugs []int
	walk []string
}

func sol(cap []int, target int) ([]string, bool) {

	return nil, false
}

func main() {
	cap := []int{4,3}
	target := 2

	walk, isPossible := sol(cap, target)

	if isPossible {
		fmt.Printf("===============================================================\n")
		fmt.Printf("=========== The Water Jug Problem – Count Min Steps ===========\n")
		fmt.Printf("===============================================================\n")
		fmt.Printf("Size of Big Jug: %v\nSize of Small Jug: %v\nTarger: %v\n\n", cap[0], cap[1], target)


		initJugs := make([]int, len(cap))

		for i, step := range walk {
			fmt.Printf("Step %v: %v\n", i+1, step)

		}
	}

}
