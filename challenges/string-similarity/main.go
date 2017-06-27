package main

import "fmt"
import "os"

func split(s string) (string, string) {
	return s[0 : len(s)/2], s[len(s)/2:]
}

type PrefixCmp struct {
	s string
	p string
}

type Queue struct {
	queue []PrefixCmp
}

func (q *Queue) Push(s PrefixCmp) {
	if len(s.s) > 0 {
		q.queue = append(q.queue, s)
	}
}

func (q *Queue) Shift() (s PrefixCmp) {
	if len(q.queue) > 0 {
		s = q.queue[0]
		q.queue = q.queue[1:]
	}
	return
}

func (q *Queue) Len() int {
	return len(q.queue)
}

func compare(s string, p string) int {
	queue := &Queue{}
	queue.Push(PrefixCmp{s[0:len(p)], p})

	sum := 0
	for queue.Len() > 0 {
		s := queue.Shift()
		if s.s == s.p {
			sum += len(s.s)
		} else if len(s.s) > 1 {
			lhs, rhs := split(s.s)
			lhp, rhp := split(s.p)

			if lhs == lhp {
				sum += len(lhs)
				queue.Push(PrefixCmp{rhs, rhp})
			} else {
				queue.Push(PrefixCmp{lhs, lhp})
			}
		}
	}

	return sum
}

func main() {
	var t int
	var s string
	fmt.Fscanf(os.Stdin, "%d\n", &t)
	for i := 0; i < t; i++ {
		sum := 0
		fmt.Fscanf(os.Stdin, "%s\n", &s)
		for j := 0; j < len(s); j++ {
			length := compare(s, s[j:])
			if l := compare(s[j:], s[j+1:]); length > l {
				j += l
				sum += (l * (l + 1)) / 2
			}
			sum += length
		}
		fmt.Printf("%d\n", sum)
	}
}
