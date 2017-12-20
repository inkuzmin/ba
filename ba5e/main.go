// Package provides solution to the
// Find a Highest-Scoring Alignment of Two Strings
// <http://rosalind.info/problems/ba5e/>
package main

import (
	"fmt"
	"github.com/inkuzmin/bautils"
	"os"
	"bufio"
	"strconv"
	"flag"
)

type Direction int

const (
	down Direction = iota
	right
	diag
)


const Inf = int(0x7FF0000000000000)
const debug = true // not used

func parseScores(file string) map[string]map[string]int {
	scores := make(map[string]map[string]int)

	f, err := os.Open(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	i := 0
	j := -1
	letters := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.Atoi(line)
		if err != nil {
			_, ok := scores[line]
			if ok {
				//lim = i
				i = 0
				j++
				continue
			} else {
				letters = append(letters, line)
				scores[line] = make(map[string]int)
			}
		} else {
			letterH := letters[j]
			letterV := letters[i]
			scores[letterH][letterV] = val
		}

		i++
	}

	return scores
}

func main() {
	pam_file := flag.String("p", "", "PAM file")
	input_file := flag.String("i", "", "Input file")

	flag.Parse()

	var v string
	var w string

	f, err := os.Open(*input_file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	l := 0
	for scanner.Scan() {
		if l == 0 {
			v = scanner.Text()
		}
		if l == 1 {
			w = scanner.Text()
		}
		l++
	}

	scores := parseScores(*pam_file)
	penalty := 5


	score, backtrack := lcs(v, w, scores, penalty)
	fmt.Printf("%v\n", score)


	cmds := outputLcs(backtrack, v, len(v) - 1, len(w) - 1)

	fmt.Printf("%v\n", cmds)

	j := 0
	for i :=0; i < len(cmds); i++ {
		switch cmd := string(cmds[i]); cmd {
		case "*":
			fmt.Printf("%c", v[j])
		case "+":
			fmt.Printf("%c", v[j])
			//j--
		case "-":
			fmt.Printf("%v", "-")
			j--
		}
		j++
	}
	fmt.Println()

	j = 0
	for i := 0; i < len(cmds); i++ {
		switch cmd := string(cmds[i]); cmd {
		case "*":
			fmt.Printf("%c", w[j])
		case "+":
			fmt.Printf("%v", "-")
			j--
		case "-":
			fmt.Printf("%c", w[j])
		}
		j++
	}
	fmt.Println()

}

func outputLcs(backtrack [][]Direction, v string, i int, j int) string {

	var helper func(int,int)string

	helper = func (i int, j int) string {
		if j == 0 && i == 0 { // this set of adhocs could be simplified
			return "*"
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
			return helper(i - 1, j) + "+"
		} else if backtrack[i][j] == right {
			return helper(i, j - 1) + "-"
		} else {
			return helper(i - 1, j - 1) + "*"
		}
	}

	return helper(i, j) // this could be rewritten as a loop with an iterator
}

func lcs(v string, w string, scores map[string]map[string]int, penalty int) (int, [][]Direction) {
	// here only two the most recent lines are necessary to store
	// or more in the case of concurrent calculations
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
		s[i][0] = -penalty * i
	}

	for j := 0; j <= len(w); j++ {
		s[0][j] = -penalty * j
	}
	s[0][0] = 0

	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {

			p := scores[string(v[i - 1])][string(w[j - 1])]


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

	return s[len(v)][len(w)], backtrack
}
