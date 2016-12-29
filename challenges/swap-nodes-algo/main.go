package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Node struct {
	data     int
	rhs, lhs *Node
}

func newNode(data int) (node *Node) {
	if data > -1 {
		node = &Node{data, nil, nil}
	}
	return node
}

func (n *Node) Swap(level int) {
	if level == 1 {
		temp := n.rhs
		n.rhs = n.lhs
		n.lhs = temp
	} else {
		if n.lhs != nil {
			n.lhs.Swap(level - 1)
		}

		if n.rhs != nil {
			n.rhs.Swap(level - 1)
		}
	}
}

func (n *Node) SetLhs(data int) *Node {
	n.lhs = newNode(data)
	return n.lhs
}

func (n *Node) SetRhs(data int) *Node {
	n.rhs = newNode(data)
	return n.rhs
}

func (n *Node) Height() int {
	height := 1
	if n.lhs != nil {
		height += n.lhs.Height()
	}

	if n.rhs != nil {
		rhs := n.rhs.Height() + 1
		if rhs > height {
			height = rhs
		}
	}
	return height
}

func (n *Node) String() string {
	buffer := &bytes.Buffer{}
	if n.lhs != nil && n.lhs.data != -1 {
		buffer.WriteString(n.lhs.String())
		buffer.WriteString(" ")
	}

	buffer.WriteString(fmt.Sprintf("%d", n.data))

	if n.rhs != nil && n.rhs.data != -1 {
		buffer.WriteString(" ")
		buffer.WriteString(n.rhs.String())
	}

	return buffer.String()
}

type Queue struct {
	q []*Node
}

func (q *Queue) Shift() (node *Node) {
	if len(q.q) > 0 {
		node = q.q[0]
		q.q = q.q[1:]
	}
	return node
}

func (q *Queue) Push(node *Node) {
	q.q = append(q.q, node)
}

func buildTree(reader io.Reader) *Node {
	q := &Queue{}
	rootNode := &Node{1, nil, nil}
	node := rootNode

	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanf(reader, "%v %v\n", &a, &b)
		newNode := node.SetLhs(a)
		if newNode != nil {
			q.Push(newNode)
		}

		newNode = node.SetRhs(b)
		if newNode != nil {
			q.Push(newNode)
		}
		node = q.Shift()
	}

	return rootNode
}

func run(reader io.Reader, writer io.Writer) {
	rootNode := buildTree(reader)
	height := rootNode.Height()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var l int
		fmt.Fscanf(reader, "%v\n", &l)
		for h := 1; h <= height; h++ {
			rootNode.Swap(l * h)
		}
		writer.Write([]byte(rootNode.String() + "\n"))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
