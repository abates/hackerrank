package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	max := 0
	counts := make(map[int]int)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if value > max {
			max = value
		}
		counts[value]++
	}
	fmt.Printf("%d\n", counts[max])
}
