package main

import (
		"testing"
		"hackerrank/tries-contacts/trie"
)

func TestTrieAdd(t *testing.T){
	tr := trie.New()
	n := tr.Add("foo")
	if tr.Size != 4 {
		t.Errorf("Expected 4, but got %d",tr.Size)
	}

	if n.Value() != 0 {
		t.Errorf("Expected 0, but got %d",n.Value)
	}
}

func TestMapKeyDoesNotExist(t *testing.T){
	testMap := make(map[int]string)
	testMap[0] = "blah"
	if _,ok := testMap[1]; ok {
		t.Errorf("Test Failed")
	}
}

func TestCallRoot(t *testing.T) {
	tr := trie.New()
	tr.Add("hack")

	root := tr.Root()
	if root == nil {
		t.Errorf("Root should not be nil")
	}
}

func TestTrieFind(t *testing.T){
	tr := trie.New()
	tr.Add("foo")
	_,ok := tr.Find("foo")
	if !ok {
		t.Fatal("Could not find node.")
	}
}

func TestCountChildrenOfNode(t *testing.T){
	tr := trie.New()
	tr.Add("hack")
	tr.Add("hackerrank")
	count := tr.FindForCountingChildren("hac")
	if count != 2 {
		t.Errorf("Expected 2, Actual %d",count)
	}
	count = tr.FindForCountingChildren("hak")
	if count != 0 {
		t.Errorf("Expected 0, Actual %d",count)
	}
}
