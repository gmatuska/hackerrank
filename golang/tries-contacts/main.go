package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	tr := NewDynamicTrie()
	reader := bufio.NewReader(os.Stdin)
	cmds_string, err := ReadLine(reader)
	Check(err)
	numCmds, err := strconv.Atoi(cmds_string)
	var results []int
	for i := 0; i < numCmds; i++ {
		cmd_string, err := ReadLine(reader)
		Check(err)
		cmds := strings.Split(cmd_string, " ")
		if cmds[0] == "add" {
			tr.Add(cmds[1])
		} else if cmds[0] == "find" {
			result := tr.Find(cmds[1])
			results = append(results, result)
		} else {
			panic("Command is not allowed")
		}
	}
	for _, i := range results {
		fmt.Println(i)
	}
}

type TrieNode struct {
	char       rune
	children   []*TrieNode
	wordsCount int
}

func NewTrieNode(r rune) *TrieNode {
	return &TrieNode{
		char:       r,
		children:   make([]*TrieNode, 0),
		wordsCount: 0,
	}
}

func (tn *TrieNode) GetChild(char rune) *TrieNode {
	for _, c := range tn.Children() {
		if c.Char() == char {
			return c
		}
	}
	return nil
}

func (tn *TrieNode) Char() rune {
	return tn.char
}

func (tn *TrieNode) Children() []*TrieNode {
	return tn.children
}

type DynamicTrie struct {
	root *TrieNode
}

func (dt *DynamicTrie) Root() *TrieNode {
	return dt.root
}

func NewDynamicTrie() *DynamicTrie {
	return &DynamicTrie{
		root: NewTrieNode('*'),
	}
}

func (dt *DynamicTrie) Add(word string) {
	curr := dt.Root()
	runes := []rune(word)
	for _, c := range runes {
		next_node := curr.GetChild(c)
		if next_node == nil {
			next_node = NewTrieNode(c)
			curr.children = append(curr.children, next_node)
		}
		next_node.wordsCount++
		curr = next_node
	}
}

func (dt *DynamicTrie) Find(prefix string) int {
	curr := dt.Root()
	runes := []rune(prefix)
	for _, c := range runes {
		next_node := curr.GetChild(c)
		if next_node == nil {
			return 0
		}
		curr = next_node
	}
	return curr.wordsCount
}

func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
