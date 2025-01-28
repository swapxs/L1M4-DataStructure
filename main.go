/*
* Solution algorithm:
* Step 1: Empty all the Jugs
* Step 2: Fill the smaller Jug(3)
* Step 3: Transfer water from Jug(3) to Jug(4), making
*         Jug(3) -> 0 and Jug(4) -> 3 (with 1 empty space)
* Step 4: Fill Jug(3)
* Step 5: Transfer water from Jug(3) to Jug(4), making
* 		  Jug(3) -> 2 and Jug(4) -> 3 + 1 (1 unit is filled
* 		  since 3 already exists)
* Step 6: Empty the Jug(4)
* Step 7: Transfer the remaining 2 units of water Jug(3) -> Jug(4)
* */

package main

import (
	"fmt"
)

func sol(cap1, cap2, t int) { }

func main() {
	bigJugCapacity := 4
	smallJugCapacity := 3
	bigJugTarget := 2

	sol(bigJugCapacity, smallJugCapacity, bigJugTarget)
}
