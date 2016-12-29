package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNodeSet(t *testing.T) {
	node := newNode(1)
	node.SetLhs(2)
	if node.lhs == nil {
		t.Errorf("Expected lhs to be non-nil")
	}

	if node.lhs.data != 2 {
		t.Errorf("Expected lhs data to be 2, got %d", node.lhs.data)
	}

	node.SetRhs(3)
	if node.rhs == nil {
		t.Errorf("Expected rhs to be non-nil")
	}

	if node.rhs.data != 3 {
		t.Errorf("Expected rhs data to be 3, got %d", node.rhs.data)
	}
}

func TestNodeString(t *testing.T) {
	node := newNode(1)
	node.SetLhs(2)
	node.SetRhs(3)
	if node.String() != "2 1 3" {
		t.Errorf("Expected '2 1 3' got '%s'", node.String())
	}

	node.lhs.SetLhs(4)
	node.rhs.SetLhs(5)
	if node.String() != "4 2 1 5 3" {
		t.Errorf("Expected '4 2 1 5 3' got '%s'", node.String())
	}

}

func TestQueuePush(t *testing.T) {
	q := &Queue{}
	if len(q.q) != 0 {
		t.Errorf("Expected length of 0, got %d", len(q.q))
	}

	q.Push(newNode(42))
	if len(q.q) != 1 {
		t.Errorf("Expected length of 1, got %d", len(q.q))
	}

	if q.q[0].data != 42 {
		t.Errorf("Expected 42, got %d", q.q[0].data)
	}
}

func TestQueueShift(t *testing.T) {
	q := &Queue{}
	q.Push(newNode(42))
	q.Push(newNode(314))

	node := q.Shift()
	if node.data != 42 {
		t.Errorf("Expected 42, got %d", node.data)
	}

	node = q.Shift()
	if node.data != 314 {
		t.Errorf("Expected 324, got %d", node.data)
	}
}

var tests = []struct {
	input      string
	level      int
	output     string
	outputSwap string
}{
	{"3\n2 3\n-1 -1\n-1 -1\n", 1, "2 1 3", "3 1 2"},
	{"5\n2 3\n-1 4\n-1 5\n-1 -1\n-1 -1\n", 2, "2 4 1 3 5", "4 2 1 5 3"},
}

func TestBuildTree(t *testing.T) {
	for i, test := range tests {
		rootNode := buildTree(strings.NewReader(test.input))
		if rootNode.String() != test.output {
			t.Errorf("Test %d expected '%s' got '%s'", i, test.output, rootNode.String())
		}
	}
}

func TestSwap(t *testing.T) {
	for i, test := range tests {
		rootNode := buildTree(strings.NewReader(test.input))
		rootNode.Swap(test.level)
		if rootNode.String() != test.outputSwap {
			t.Errorf("Test %d expected '%s' got '%s'", i, test.outputSwap, rootNode.String())
		}
	}
}

func TestRun(t *testing.T) {
	files, err := filepath.Glob("testdata/test*_input.txt")
	if err != nil {
		panic(err.Error())
	}

	for _, filename := range files {
		// read the expected output first
		outputFilename := strings.TrimSuffix(filename, "_input.txt") + "_output.txt"
		expectedResult, err := ioutil.ReadFile(outputFilename)
		if err != nil {
			panic(err.Error())
		}

		file, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}

		writer := &bytes.Buffer{}
		run(file, writer)

		result := writer.String()
		if result != string(expectedResult) {
			t.Errorf("Test %s expected '%s' got' %s'", filename, string(expectedResult), result)
		}
	}
}
