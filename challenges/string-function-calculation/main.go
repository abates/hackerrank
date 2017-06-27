package main

import "fmt"
import "os"
import "sort"

func f(s string) int {
	count := 0
	return len(s) * count
}

func main() {
	s := ""
	fmt.Fscanf(os.Stdin, "%s\n", &s)
	counts := make([]int, 1)
	counts[0] = 1
	substrings := make(map[string]bool)

	for start := 0; start < len(s); start++ {
		for end := start + 1; end <= len(s); end++ {
			substring := s[start:end]
			if _, ok := substrings[substring]; !ok {
				substrings[substring] = true
				occurances := 0
				for i := 0; i <= len(s)-len(substring); i++ {
					if substring == s[i:i+len(substring)] {
						occurances++
					}
				}
				counts = append(counts, len(substring)*occurances)
				println("Next count", counts[len(counts)-1])
			}
		}
	}
	sort.Ints(counts)
	fmt.Fprintf(os.Stderr, "%v\n", counts)
	fmt.Printf("%d\n", counts[len(counts)-1])
}
