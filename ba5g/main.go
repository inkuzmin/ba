// Package provides solution to the
// Compute the Edit Distance Between Two Strings
// <http://rosalind.info/problems/ba5g/>
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
	match
	mismatch
)

func main() {
	input_file := flag.String("i", "", "Input file")
	flag.Parse()

	v, w := bautils.ParseSequences(*input_file)

	_, backtrack := lcs(v, w)

	var d = 0
	outputLcs(backtrack, v, len(v) - 1, len(w) - 1, &d)

	fmt.Printf("\nLevenshtein distance: %v\n", d)
}

func outputLcs(backtrack [][]Direction, v string, i int, j int, d *int) {
	if i < 0 || j < 0 {
		for j > 0 {
			*d += 1
			j--
		}
		for i > 0 {
			*d += 1
			i--
		}
		return
	} else if backtrack[i][j] == down {
		*d += 1
		outputLcs(backtrack, v, i - 1, j, d)
	} else if backtrack[i][j] == right {
		*d += 1
		outputLcs(backtrack, v, i, j - 1, d)
	} else if backtrack[i][j] == mismatch {
		*d += 1
		outputLcs(backtrack, v, i - 1, j - 1, d)
	} else {
		outputLcs(backtrack, v, i - 1, j - 1, d)
	}
	return
}

func lcs(v string, w string) (int, [][]Direction) {
	penalty := 1

	s := make([][]int, len(v)+1)
	for i := range s {
		s[i] = make([]int, len(w)+1)
	}

	backtrack := make([][]Direction, len(v))
	for i := range backtrack {
		backtrack[i] = make([]Direction, len(w))
	}

	for i := 0; i <= len(v); i++ {
		s[i][0] = -penalty * i
	}

	for j := 0; j <= len(w); j++ {
		s[0][j] = -penalty * j
	}
	s[0][0] = 0

	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			p := -1
			if v[i-1] == w[j-1] {
				p = 0
			}

			s[i][j] = bautils.Max3(s[i-1][j]-penalty, s[i][j-1]-penalty, s[i-1][j-1]+p)

			var d Direction
			if s[i][j] == s[i-1][j-1] - 1 {
				d = mismatch
		    } else if s[i][j] == s[i-1][j] - penalty{
				d = down
			} else if s[i][j] == s[i][j-1] - penalty{
				d = right
			} else if s[i][j] == s[i-1][j-1] {
				d = match
			}
			backtrack[i-1][j-1] = d
		}
	}

	return s[len(v)][len(w)], backtrack
}
