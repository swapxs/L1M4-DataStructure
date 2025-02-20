package main

import "fmt"

type State struct {
	mOnLeft, cOnLeft, mOnRight, cOnRight int
	boat string
}

func isValid(s State) bool {
	if s.mOnLeft < 0 || s.cOnLeft < 0 || s.mOnRight < 0 || s.cOnRight < 0 {
		return false
	}

	if s.mOnLeft > 0 && s.mOnLeft < s.cOnLeft {
		return false
	}

	if s.mOnRight > 0 && s.mOnRight < s.cOnRight {
		return false
	}

	return true
}

func main() {
	initState := State{3, 3, 0, 0, "left"}
	goal := State{0, 0, 3, 3, "right"}

	q := []State{initState}

	visited := make(map[State]bool)
	visited[initState] = true

	pNode := make(map[State]State)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == goal {
			walk := []State{curr}
			for curr != initState {
				curr = pNode[curr]
				walk = append([]State{curr}, walk...)
			}

			for i := 0; i < len(walk)-1; i++ {
				mBoat := 0
				cBoat := 0

				if walk[i].boat == "left" {
					mBoat = walk[i+1].mOnRight - walk[i].mOnRight
					cBoat = walk[i+1].cOnRight - walk[i].cOnRight
				} else {
					mBoat = walk[i].mOnRight - walk[i+1].mOnRight
					cBoat = walk[i].cOnRight - walk[i+1].cOnRight
				}

				fmt.Printf("(%v, %v)", walk[i].mOnLeft, walk[i].cOnLeft)

				if walk[i].boat == "left" {
					fmt.Printf("\t\t->(%v,%v)->\t(%v, %v)\n", mBoat, cBoat, walk[i+1].mOnRight, walk[i+1].cOnRight)
				} else {
					fmt.Printf("\t\t<-(%v,%v)<-\t(%v, %v)\n", mBoat, cBoat, walk[i+1].mOnRight, walk[i+1].cOnRight)
				}
			}
			return
		}

		mvs := [][2]int{
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 1},
			{2, 0},
		}

		for _, move := range mvs {
			m, c := move[0], move[1]
			nxt := curr

			if curr.boat == "left" {
				nxt.mOnLeft -= m
				nxt.cOnLeft -= c
				nxt.mOnRight += m
				nxt.cOnRight += c
				nxt.boat = "right"

			} else {
				nxt.mOnLeft += m
				nxt.cOnLeft += c
				nxt.mOnRight -= m
				nxt.cOnRight -= c
				nxt.boat = "left"
			}

			if isValid(nxt) && !visited[nxt] {
				visited[nxt] = true
				pNode[nxt] = curr
				q = append(q, nxt)
			}
		}
	}
}
