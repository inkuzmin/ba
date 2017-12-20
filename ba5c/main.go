// Package provides solution to the
// Find a Longest Common Subsequence of Two Strings
// <http://rosalind.info/problems/ba5c/>
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
)


const Inf = int(0x7FF0000000000000)


func main() {
	input_file := flag.String("i", "", "Input file")
	flag.Parse()

	v, w := bautils.ParseSequences(*input_file)

	score, backtrack := lcs(v, w)

	fmt.Printf("%v\n", score)
	outputLcs(backtrack, v, len(v) - 1, len(w) - 1)
	fmt.Println()
}

func outputLcs(backtrack [][]Direction, v string, i int, j int) {
	if i < 0 || j < 0 {
		return
	} else if backtrack[i][j] == down {
		outputLcs(backtrack, v, i - 1, j)
	} else if backtrack[i][j] == right {
		outputLcs(backtrack, v, i, j - 1)
	} else {
		outputLcs(backtrack, v, i - 1, j - 1)
		fmt.Printf("%c", v[i])
	}
}

func lcs(v string, w string) (int, [][]Direction) {
	s := make([][]int, len(v)+1)
	for i := range s {
		s[i] = make([]int, len(w)+1)
	}

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
	s[0][0] = 0

	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			p := 0
			if v[i - 1] == w[j - 1] {
				p = 1
			}

			s[i][j] = bautils.Max3(s[i - 1][j], s[i][j - 1], s[i - 1][j - 1] + p)

			var d Direction
			if s[i][j] == s[i-1][j] {
				d = down
			} else if s[i][j] == s[i][j-1] {
				d = right
			} else if s[i][j] == (s[i-1][j-1] + 1) {
				d = diag
			}
			backtrack[i-1][j-1] = d
		}
	}

	return s[len(v)][len(w)], backtrack
}
