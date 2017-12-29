package bautils

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

const Inf = int(0x7FF0000000000000)


func Max(vs ...int) (max int) {
	max = -Inf
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return
}

func Max2(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Max3(a int, b int, c int) int {
	return Max(Max(a, b), c)
}

func Max4(a int, b int, c int, d int) int {
	return Max(Max(Max(a, b), c), d)
}

// Thanks peterSO for an efficient runic solution
// https://stackoverflow.com/a/4966500
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

func ParseSequences(file string) (string, string) {
	var v string
	var w string

	f, err := os.Open(file)

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

	return v, w
}

func ParsePAM(file string) map[string]map[string]int {
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