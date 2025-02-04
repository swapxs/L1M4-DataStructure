package main

import (
	"fmt"
	"strconv"
)

type State struct {
	jug1, jug2 int
	walk []string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sol(cap1, cap2, target int) {
	initState := State{0,0, []string{}}

	q := []State{initState}

	visited := make(map[string]bool)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.jug1 == target {
			fmt.Printf("Achieved the solution in %v steps\n", len(curr.walk))
			for i, step := range curr.walk {
				fmt.Printf("Step %v: %v\n", i+1, step)
			}
			return
		}

		node := strconv.Itoa(curr.jug1) + "," + strconv.Itoa(curr.jug2)

		if visited[node] {
			continue
		}

		visited[node] = true

		// Fill jug1
		if curr.jug1 < cap1 {
			nxt := State{
				cap1,
				curr.jug2, append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Fill Jug 1\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))
			q = append(q, nxt)
		}

		// Fill jug2
		if curr.jug2 < cap2 {
			nxt := State{
				curr.jug1,
				cap2,
				append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Fill Jug 2\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))
			q = append(q, nxt)
		}

		// Empty jug1
		if curr.jug1 > 0 {
			nxt := State{
				0,
				curr.jug2,
				append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Empty Jug 1\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))

			q = append(q, nxt)

		}

		// Empty jug2
		if curr.jug2 > 0 {
			nxt := State{
				curr.jug1,
				0,
				append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Empty Jug 2\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))
			q = append(q, nxt)
		}

		// Pour jug1 to jug2
		if curr.jug1 > 0 && curr.jug2 < cap2 {
			pour := min(curr.jug1, cap2-curr.jug2)

			nxt := State{
				curr.jug1 - pour,
				curr.jug2 + pour,
				append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Pour Jug 1 to Jug 2\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))
			q = append(q, nxt)
		}

		// Pour jug2 to jug1
		if curr.jug2 > 0 && curr.jug1 < cap1 {
			pour := min(curr.jug2, cap1-curr.jug1)

			nxt := State{
				curr.jug1 + pour,
				curr.jug2 - pour,
				append([]string{}, curr.walk...),
			}

			nxt.walk = append(nxt.walk, fmt.Sprintf("Pour Jug 2 to Jug 1\n\tJug 1: %v, Jug 2: %v", nxt.jug1, nxt.jug2))
			q = append(q, nxt)
		}

	}
	fmt.Printf("No solution found.")
}

func main() {

	jug1cap := 4
	jug2cap := 3
	trgt := 2

	fmt.Printf("Size of Jug 1: %v\nSize of Jug 2: %v\nTarget: %v\n", jug1cap, jug2cap, trgt)
	sol(jug1cap, jug2cap, trgt)

}
