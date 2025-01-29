// Implementating the program using BFS
package main

import (
	"container/list"
	"fmt"
)

type State struct {
	jugs []int
	walk []string
}

func toString(s []int) string {
    var str string
    for _, val := range s {
        str += fmt.Sprintf("%d-", val)
    }
    return str
}

func contains(x []int, t int) bool {
	for _, val := range x {
		if val == t {
			return true
		}
	}
	return false
}

func sol(cap []int, target int) ([]string, bool) {
	initState := State{
		jugs: make([]int, len(cap)),
		walk: []string{},
	}

	q := list.New()
	q.PushBack(initState)

	visited := make(map[string]bool)
	visited[toString(initState.jugs)] = true

	for q.Len() > 0{
		item := q.Front()
		q.Remove(item)
		curr := item.Value.(State)

		if contains(curr.jugs, target) {
			return curr.walk, true
		}

		// Operation: Fill Jug
		for i := 0; i < len(cap); i++ {
			if curr.jugs[i] < cap[i] {
				nxt := make([]int, len(curr.jugs))
				copy(nxt, curr.jugs)

				nxt[i] = cap[i]

				newState := State{
					jugs: nxt,
					walk: append(curr.walk, fmt.Sprintf("Fill Jug %v to %v gallons", i+1, cap[i])),
				}

				node := toString(newState.jugs)

				if !visited[node] {
					visited[node] = true
					q.PushBack(newState)
				}
			}
		}

		// Operation: Empty Jug
		for i := 0; i < len(cap); i++ {
			if curr.jugs[i] > 0 {
				nxt := make([]int, len(curr.jugs))
				copy(nxt, curr.jugs)

				nxt[i] = 0

				newState := State{
					jugs: nxt,
					walk: append(curr.walk, fmt.Sprintf("Empty Jug %v", i+1)),
				}

				node := toString(newState.jugs)

				if !visited[node] {
					visited[node] = true
					q.PushBack(newState)
				}
			}
		}

		// Operation: Pour
		for i := 0; i < len(cap); i++ {
			if curr.jugs[i] > 0 {
				for j := 0; j < len(cap); j++ {
					if i != j {
						pour := min(curr.jugs[i], cap[j] - curr.jugs[j])
						nxt := make([]int, len(curr.jugs))

						copy(nxt, curr.jugs)

						nxt[i] -= pour
						nxt[j] += pour

						newState := State{
							jugs: nxt,
							walk: append(curr.walk, fmt.Sprintf("Pour Jug %v into Jug %v", i+1, j+1)),
						}

						node := toString(newState.jugs)

						if !visited[node] {
							visited[node] = true
							q.PushBack(newState)
						}
					}
				}
			}
		}
	}

	return nil, false
}

func main() {
	cap := make([]int, 2)

	// ------- DATA -------
	cap[0] = 4
	cap[1] = 3
	target := 2
	// --------------------

	walk, isPossible := sol(cap, target)

	if isPossible {
		fmt.Printf("===============================================================\n")
		fmt.Printf("=========== The Water Jug Problem – Count Min Steps ===========\n")
		fmt.Printf("===============================================================\n")
		fmt.Printf("Size of Big Jug: %v\nSize of Small Jug: %v\nTarger: %v\n\n", cap[0], cap[1], target)


		initJugs := make([]int, len(cap))

		for i, step := range walk {
			fmt.Printf("Step %v: %v\n", i+1, step)

			if step[:4] == "Fill" {
				var jugNo, gallons int
				fmt.Sscanf(step, "Fill Jug %v to %v gallons", &jugNo, &gallons)
				initJugs[jugNo - 1] = gallons
			} else if step[:5] =="Emty" {
				var jugNo int
				fmt.Sscanf(step, "Empty Jug %v ", &jugNo)
				initJugs[jugNo - 1] = 0
			} else if step[:4] == "Pour" {
				var src, dest int
				fmt.Sscanf(step, "Pour Jug %v into Jug %v", &src, &dest)

				pourAmount := min(initJugs[src - 1], cap[dest - 1] - initJugs[dest - 1])
				initJugs[src - 1] -= pourAmount
				initJugs[dest - 1] += pourAmount
			}

			fmt.Printf("\tWater in Big Jug: %v, Water in Small Jug: %v\n\n", initJugs[0], initJugs[1])
		}
	} else {
		fmt.Printf("No Solution.")
	}

}
