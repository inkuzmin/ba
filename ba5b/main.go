// Package provides solution to the
// Find the Length of a Longest Path in a Manhattan-like Grid
// <http://rosalind.info/problems/ba5b/>
package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"strconv"
	"os"
)

func main() {
	file := flag.String("f", "", "Input file")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		panic(err)
	}

	data := string(bytes)

	var n int
	var m int
	var down [][]int
	var right [][]int

	var separator bool
	separator = false

	x := 0 // indexes for filling down and right
	y := 0 // indexes for filling down and right

	for i, ln := range strings.Split(data, "\n") {
		for j, c := range strings.Split(ln, " ") {
			if i == 0 {
				if j == 0 {
					n, err = strconv.Atoi(c)
					if err != nil {
						fmt.Fprintf(os.Stderr, "error: %v\n", err)
						os.Exit(1)
					}
				}
				if j == 1 {
					m, err = strconv.Atoi(c)
					if err != nil {
						fmt.Fprintf(os.Stderr, "error: %v\n", err)
						os.Exit(1)
					}
				}
			} else if i > 0 {
				if i == 1 && j == 0 {
					down = make([][]int, n)
					for i := range down {
						down[i] = make([]int, m + 1)
					}

					right = make([][]int, n + 1)
					for i := range right {
						right[i] = make([]int, m)
					}
				}

				if ln == "-" {
					separator = true
					x = 0
					y = 0

					continue
				}

				if !separator {
					v, err := strconv.Atoi(c)
					if err != nil {
						fmt.Fprintf(os.Stderr, "error: %v\n", err)
						os.Exit(1)
					}
					down[x][y] = v
					y++
					if y > m {
						x++
						y = 0
					}
				} else {
					v, err := strconv.Atoi(c)
					if err != nil {
						fmt.Fprintf(os.Stderr, "error: %v\n", err)
						os.Exit(1)
					}
					right[x][y] = v
					y++
					if y > (m - 1) {
						y = 0
						x++
					}
				}
			}
		}
	}

	s := manhattanTourist(n, m, down, right)

	fmt.Printf("s[%v][%v] = %v\ns = %v", n, m, s[n][m], s)
}

func manhattanTourist(n int, m int, down[][]int, right [][]int) [][]int {
	s := make([][]int, n + 1)
	for i := range s {
		s[i] = make([]int, m + 1)
	}

	s[0][0] = 0

	for i := 1; i <= n; i++ {
		s[i][0] = s[i - 1][0] + down[i-1][0]
	}
	for j := 1; j <= m; j++ {
		s[0][j] = s[0][j - 1] + right[0][j-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = max(s[i-1][j] + down[i-1][j], s[i][j-1] + right[i][j-1])
		}
	}
	return s
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}