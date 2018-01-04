// Package provides solution to the
// Find a Highest-Scoring Local Alignment of Two Strings
// <http://rosalind.info/problems/ba5f/>
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
	pam_file := flag.String("p", "", "PAM file")
	input_file := flag.String("i", "", "Input file")

	flag.Parse()

	scores := bautils.ParsePAM(*pam_file)
	v, w := bautils.ParseSequences(*input_file)

	score, backtrack, ipenult, jpenult := lcs(v, w, scores, 5)

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




func lcs(v string, w string, scores map[string]map[string]int, penalty int) (int, [][]Direction, int, int) {

	s := make([][]int, len(v)+1)

	for i := range s {
		s[i] = make([]int, len(w)+1)
	}

	s[0][0] = 0

	backtrack := make([][]Direction, len(v))
	for i := range backtrack {
		backtrack[i] = make([]Direction, len(w))
	}

	for i := 0; i <= len(v); i++ {
		s[i][0] = 0
	}

	for j := 0; j <= len(w); j++ {
		s[0][j] = 0
	}

	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			p := scores[string(v[i - 1])][string(w[j - 1])]

			s[i][j] = bautils.Max4(0, s[i-1][j] - penalty, s[i][j-1] - penalty, s[i-1][j-1] + p)

			var d Direction
			if s[i][j] == s[i-1][j] - penalty {
				d = down
			} else if s[i][j] == s[i][j-1] - penalty {
				d = right
			} else if s[i][j] == (s[i-1][j-1] + p) {
				d = diag
			} else if s[i][j] == 0 {
				d = taxi
			}

			backtrack[i-1][j-1] = d
		}
	}

	smax := 0
	ipenult := 0
	jpenult := 0

	for i := 0; i <= len(v); i++ {
		for j := 0; j <= len(w); j++ {
			if s[i][j] > smax {
				smax = s[i][j]
				ipenult = i
				jpenult = j
			}
		}
	}

	return smax, backtrack, ipenult, jpenult
}

