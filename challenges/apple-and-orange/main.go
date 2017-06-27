package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scanNums(scanner *bufio.Scanner, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	return nums
}

func calculate(distances []int, start int, s int, t int) int {
	result := 0
	for _, distance := range distances {
		delta := start + distance
		if delta >= s && delta <= t {
			result++
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	st := scanNums(scanner, 2)
	ab := scanNums(scanner, 2)
	mn := scanNums(scanner, 2)
	appleDistances := scanNums(scanner, mn[0])
	orangeDistances := scanNums(scanner, mn[1])

	numApples := calculate(appleDistances, ab[0], st[0], st[1])
	numOranges := calculate(orangeDistances, ab[1], st[0], st[1])

	fmt.Printf("%d\n", numApples)
	fmt.Printf("%d\n", numOranges)
}
