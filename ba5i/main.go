// Package provides solution to the
// Find a Highest-Scoring Overlap Alignment of Two Strings
// <http://rosalind.info/problems/ba5i/>
package main

import (
	"fmt"
	"flag"
	"github.com/inkuzmin/bautils"
)

type Direction int

const (
	down Direction = iota
	right
	diag
	taxi
)

const Inf = int(0x7FF0000000000000)

func main() {
	input_file := flag.String("i", "", "Input file")
	flag.Parse()

	v, w := bautils.ParseSequences(*input_file)

	score, backtrack, ipenult, jpenult := lcs(v, w)

	fmt.Printf("%v\n", score)

	cmds, va, wa := outputLcs(backtrack, v, w, ipenult, jpenult)

	fmt.Printf("%v\n", cmds)
	fmt.Printf("%v\n", va)
	fmt.Printf("%v\n", wa)

}

func outputLcs(backtrack [][]Direction, v string, w string, i int, j int) (string, string, string) {

	var helper func(int,int) string

	va := "" // v aligned
	wa := "" // w aligned

	// this could be rewritten as a loop with an iterator
	helper = func (i int, j int) string {
		if j == 0 && i == 0 { // this set of adhocs could be simplified
			if backtrack[i][j] == diag {
				va += string(v[i])
				wa += string(w[j])
				return "*"
			} else if backtrack[i][j] == taxi {
				return "^"
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
		if backtrack[i][j] == down {
			va += string(v[i])
			wa += "-"
			return helper(i - 1, j) + "+"
		} else if backtrack[i][j] == right {
			va += "-"
			wa += string(w[j])
			return helper(i, j - 1) + "-"
		} else if backtrack[i][j] == diag {
			va += string(v[i])
			wa += string(w[j])
			return helper(i - 1, j - 1) + "*"
		} else if backtrack[i][j] == taxi {
			return "^"
		} else {
			panic(-1)
			return "?"
		}
	}

	return helper(i - 1, j - 1), bautils.Reverse(va), bautils.Reverse(wa)
}




func lcs(v string, w string) (int, [][]Direction, int, int) {

	s := make([][]int, len(v)+1)

	for i := range s {
		s[i] = make([]int, len(w)+1)
	}

	s[0][0] = 0

	penalty := 2

	backtrack := make([][]Direction, len(v))
	for i := range backtrack {
		backtrack[i] = make([]Direction, len(w))
	}

	for i := 0; i <= len(v); i++ {
		s[i][0] = 0
	}

	for j := 0; j <= len(w); j++ {
		s[0][j] = -penalty * j
	}



	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			p := -2
			if v[i-1] == w[j-1] {
				p = 1
			}

			s[i][j] = bautils.Max3(s[i-1][j] - penalty, s[i][j-1] - penalty, s[i-1][j-1] + p)

			var d Direction
			if s[i][j] == s[i-1][j] - penalty {
				d = down
			} else if s[i][j] == s[i][j-1] - penalty {
				d = right
			} else if s[i][j] == (s[i-1][j-1] + p) {
				d = diag
			}

			backtrack[i-1][j-1] = d
		}
	}

	smax := 0
	ipenult := 0
	jpenult := 0

	i := len(v)

	for j := 0; j <= len(w); j++ {
		if s[i][j] > smax {
			smax = s[i][j]
			ipenult = i
			jpenult = j
		}
	}

	return smax, backtrack, ipenult, jpenult
}

