// Package provides solution to the
// Find the Minimum Number of Coins Needed to Make Change
// <http://rosalind.info/problems/ba5a/>
package main

import (
	"fmt"
	"flag"
	"strings"
	"strconv"
	"os"
)

const Inf = int(0x7FF0000000000000)

func main() {
	money := flag.Int("money", 0, "Amount of money to exchange")
	coinsArg := flag.String("coins", "", "Coins to exchange money")

	flag.Parse()

	coins := stringToCoins(*coinsArg)

	minNumCoins, list := dynamicChange3(*money, coins)

	fmt.Printf("Minimum number of coins is: %v, here they are: %v\n", minNumCoins, list)
}

func stringToCoins(csv string) []int {
	var coins []int
	for _, c := range strings.Split(csv, ",") {
		i, err := strconv.Atoi(c)
		if err != nil {
			os.Exit(1)
		}
		coins = append(coins, i)
	}

	return coins
}

func dynamicChange3(money int, coins []int) (int, []int) {
	size := max(coins)

	var n int

	var minNumCoins []int
	minNumCoins = make([]int, size+1)

	var coinSet [][]int
	coinSet = make([][]int, size+1)

	for m := 1; m <= money; m++ {

		if m > size {
			minNumCoins = shift(minNumCoins)
			coinSet = shift2(coinSet)
			n = size
		} else {
			n = m
		}

		minNumCoins[n] = Inf
		for _, coin := range coins {
			if m >= coin {
				if minNumCoins[n-coin]+1 < minNumCoins[n] {
					minNumCoins[n] = minNumCoins[n-coin] + 1
					coinSet[n] = append(coinSet[n-coin], coin)
				}
			}
		}

	}

	return minNumCoins[n], coinSet[n]
}

func dynamicChange2(money int, coins []int) (int, []int) {
	size := max(coins)
	var minNumCoins []int
	minNumCoins = make([]int, size+1)
	var n int
	for m := 1; m <= money; m++ {

		if m > size {
			minNumCoins = shift(minNumCoins)
			n = size
		} else {
			n = m
		}

		minNumCoins[n] = Inf

		for _, coin := range coins {
			if m >= coin {
				if minNumCoins[n-coin]+1 < minNumCoins[n] {
					minNumCoins[n] = minNumCoins[n-coin] + 1
				}
			}
		}

	}

	return minNumCoins[n], minNumCoins
}

func shift(list []int) []int {
	i := 0
	for ; i < len(list)-1; i++ {
		list[i] = list[i+1]
	}
	list[i] = Inf
	return list
}

func shift2(list [][]int) [][]int {
	i := 0
	for ; i < len(list)-1; i++ {
		list[i] = list[i+1]
	}
	list[i] = nil
	return list
}

func max(slice []int) int {
	m := -1

	for _, v := range slice {
		if v > m {
			m = v
		}
	}

	return m
}

func sumCoins(numCoins []int, coins []int) int {
	t := 0
	for i, n := range coins {
		t += n * numCoins[i]
	}
	return t
}

func dynamicChange(money int, coins []int) int {
	var minNumCoins []int
	minNumCoins = make([]int, money+1)

	for m := 1; m <= money; m++ {
		minNumCoins[m] = Inf

		for _, coin := range coins {
			if m >= coin {
				if minNumCoins[m-coin]+1 < minNumCoins[m] {
					minNumCoins[m] = minNumCoins[m-coin] + 1
				}
			}
		}
	}

	return minNumCoins[money]
}

func recursiveChange(money int, coins []int) int {
	if money == 0 {
		return 0
	}

	minNumCoins := Inf

	for _, coin := range coins {
		if money >= coin {
			numCoins := recursiveChange(money-coin, coins)
			if numCoins+1 < minNumCoins {
				minNumCoins = numCoins + 1
			}
		}
	}
	return minNumCoins
}
