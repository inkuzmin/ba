// Package provides solution to the
// Align Two Strings Using Affine Gap Penalties
// <http://rosalind.info/problems/ba5j/>
package main

import (
	"fmt"
	"flag"
	"github.com/inkuzmin/bautils"
)

type Direction int

const (
	down Direction = iota // 0
	right // 1
	diag // 2
	elevator_up // 3
	elevator_down // 4
)

func main() {
	pam_file := flag.String("p", "", "PAM file")
	input_file := flag.String("i", "", "Input file")

	flag.Parse()

	scores := bautils.ParsePAM(*pam_file)
	v, w := bautils.ParseSequences(*input_file)

	epsilon := 1
	sigma := 11

	score, backtracks := lcs(v, w, scores, sigma, epsilon)

	fmt.Printf("%v\n", score)

	cmds, va, wa := outputLcs(backtracks, v, w, len(v), len(w))

	fmt.Printf("%v\n", cmds)
	fmt.Printf("%v\n", va)
	fmt.Printf("%v\n", wa)
}


func outputLcs(backtracks [3][][]Direction, v string, w string, i int, j int) (string, string, string) {

	var helper func(int,int) string

	va := "" // v aligned
	wa := "" // w aligned

	backtrackMiddle := backtracks[0]
	backtrackLower := backtracks[1]
	backtrackUpper := backtracks[2]

	backtrack := backtrackMiddle
	currentBacktrack := "middle"

	// this could be rewritten as a loop with an iterator
	helper = func (i int, j int) string {
		if j == 0 && i == 0 { // this set of adhocs could be simplified
			if backtrack[i][j] == diag {
				va += string(v[i])
				wa += string(w[j])
				return "*"
			}
		} else if j < 0 || i < 0 {
			if j < i {
				return "+"
			} else if i < j {
				return "-"
			} else {
				return "*"
			}
		}
		if backtrack[i][j] == elevator_down {
			if currentBacktrack == "middle" {
				backtrack = backtrackLower
				currentBacktrack = "lower"
				return helper(i, j)
			} else if currentBacktrack == "upper" {
				backtrack = backtrackMiddle
				currentBacktrack = "middle"
				va += "-"
				wa += string(w[j])
				return helper(i, j - 1) + "-"
			}
		} else if backtrack[i][j] == elevator_up {

			if currentBacktrack == "lower" {
				backtrack = backtrackMiddle
				currentBacktrack = "middle"
				va += string(v[i])
				wa += "-"
				return helper(i - 1, j) + "+"
			} else if currentBacktrack == "middle" {
				backtrack = backtrackUpper
				currentBacktrack = "upper"
				return helper(i, j)
			}
		} else if backtrack[i][j] == diag {
			va += string(v[i])
			wa += string(w[j])
			return helper(i - 1, j - 1) + "*"
		} else if backtrack[i][j] == down {
			va += string(v[i])
			wa += "-"
			return helper(i - 1, j) + "+"
		} else if backtrack[i][j] == right {
			va += "-"
			wa += string(w[j])
			return helper(i, j - 1) + "-"
		}

		return "?"
	}

	return helper(i - 1, j - 1), bautils.Reverse(va), bautils.Reverse(wa)
}



func lcs(v string, w string, scores map[string]map[string]int, sigma int, epsilon int) (int, [3][][]Direction) {

	lower  := make([][]int, len(v)+1)
	middle := make([][]int, len(v)+1)
	upper  := make([][]int, len(v)+1)

	for i := range lower {
		lower[i]  = make([]int, len(w)+1)
		middle[i] = make([]int, len(w)+1)
		upper[i]  = make([]int, len(w)+1)
	}

	lower[0][0]  = 0
	middle[0][0] = 0
	upper[0][0]  = 0

	//for i := 0; i <= len(v); i++ {
	//	lower[i][0] = 0
	//}
	//
	//for j := 0; j <= len(w); j++ {
	//	upper[0][j] = 0
	//}

	backtrackMiddle := make([][]Direction, len(v))
	backtrackLower := make([][]Direction, len(v))
	backtrackUpper := make([][]Direction, len(v))
	for i := range backtrackMiddle {
		backtrackMiddle[i] = make([]Direction, len(w))
		backtrackLower[i] = make([]Direction, len(w))
		backtrackUpper[i] = make([]Direction, len(w))
	}

	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			score := scores[string(v[i-1])][string(w[j-1])]

			lower[i][j] = bautils.Max(lower[i-1][j]-epsilon, middle[i-1][j]-sigma)
			upper[i][j] = bautils.Max(upper[i][j-1]-epsilon, middle[i][j-1]-sigma)

			middle[i][j] = bautils.Max(lower[i][j], middle[i-1][j-1]+score, upper[i][j])


			// TODO: store backtrack in 1 matrix instead of 3
			var d Direction
			if middle[i][j] == middle[i-1][j-1]+score {
				d = diag
			} else if middle[i][j] == upper[i][j] {
				d = elevator_up
			} else if middle[i][j] == lower[i][j] {
				d = elevator_down
			}
			backtrackMiddle[i-1][j-1] = d

			if lower[i][j] == lower[i-1][j]-epsilon {
				d = down
			} else if lower[i][j] == middle[i-1][j]-sigma {
				d = elevator_up
			}
			backtrackLower[i-1][j-1] = d

			if upper[i][j] == upper[i][j-1]-epsilon {
				d = right
			} else if upper[i][j] == middle[i][j-1]-sigma {
				d = elevator_down
			}
			backtrackUpper[i-1][j-1] = d

		}
	}

	backtracks := [3][][]Direction{
		backtrackMiddle,
		backtrackLower,
		backtrackUpper,
	}

	return middle[len(v)][len(w)], backtracks
}
