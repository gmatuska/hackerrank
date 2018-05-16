package trie

const nul = 0x0

type Trie struct {
	root *Node
	Size int
}
type Node struct {
	value rune
	children map[rune]*Node
	parent *Node
	isCompleteWord bool
}

func (trie *Trie) Add(str string) *Node {
	runes := []rune(str)
	node := trie.root
	for i := range runes {
		r := runes[i]
		//check if rune is in children already
		if n,ok := node.children[r];ok {
			node = n
		} else {
			//add rune to children
			node = node.NewChild(r,false)
			trie.Size++
		}
	}
	node = node.NewChild(nul,true)
	trie.Size++
	return node
}

func (node *Node) NewChild(val rune,isWord bool) *Node {
	child := &Node{
		value:          val,
		parent:			node,
		children:       make(map[rune]*Node),
		isCompleteWord: isWord,
	}
	node.children[val] = child
	return child
}

func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
			isCompleteWord:false,
			value:0,
		},
	}
}

func (t *Trie) FindForCountingChildren(key string)int {
	n := findNode(t.Root(),[]rune(key))
	return n.CountChildren()
}

func (t *Trie) Find(key string) (*Node, bool) {
	//root := t.Root()
	node := findNode(t.Root(),[]rune(key))
	if node == nil {
		return nil,false
	}

	node,ok := node.Children()[nul]
	if !ok || !node.isCompleteWord {
		return node, false
	}
	return node,true
}

func (n *Node) Children()map[rune]*Node {
	return n.children
}

// Returns the root node for the Trie.
func (t *Trie) Root() *Node {
	return t.root
}

func (n *Node) Value() rune {
	return n.value
}

func findNode(node *Node, runes []rune) *Node {
	if node == nil{
		return nil
	}

	if len(runes) == 0 {
		return node
	}

	n,ok := node.Children()[runes[0]]
	if !ok {
		return nil
	}

	var nrunes []rune
	if len(runes) > 1{
		nrunes = runes[1:]
	} else{
		nrunes = runes[0:0]
	}
	return findNode(n,nrunes)
}

func (n *Node) CountChildren() int {
	count := 0
	if n == nil {
		return count
	}
	for k,_ := range n.Children(){
		//count zeros
		if k == 0{
			count++
		} else {
			count += n.Children()[k].CountChildren()
		}
	}
	return count
}
