package main

type Vertex struct {
	val   int64
	edges []int64
}

func New(val int64)*Vertex{
	return &Vertex{
		val:	val,
	}
}

type Stack struct {
	items []int64
}

func(s *Stack) Top()int64{
	return s.items[0]
}

func(s *Stack) Empty()bool{
	return len(s.items) == 0
}

func(s *Stack) Push(item int64)*Stack{
	s.items = append(s.items,item)
	return s
}

func(s *Stack) Pop()int64{
	if len(s.items) > 0{
		item := s.items[len(s.items)-1]
		s.items = s.items[1:]
		return item
	}
	return nil
}

type Graph struct {
	vertices Vertices
}

type Vertices []*Vertex

func Find(v Vertices, from int64) *Vertex {
	for _, i := range v {
		if i.val == from {
			return i
		}
	}
	return nil
}

//Check if edge exists already in edges collection
func All(e []int64,val int64) bool {
	for _, i := range e{
		if i == val{
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to int64) *Graph {
	fromVertex := Find(g.vertices, from)
	if fromVertex == nil {
		fromVertex = &Vertex{
			val: from,
		}
		g.vertices = append(g.vertices,fromVertex)
	}
	//check for duplicate edges
	if !All(fromVertex.edges,to){
		fromVertex.edges = append(fromVertex.edges,to)
	}
	toVertex := Find(g.vertices, to)
	if toVertex == nil {
		toVertex = &Vertex{
			val: to,
		}
		g.vertices = append(g.vertices,toVertex)
	}
	if !All(toVertex.edges,from){
		toVertex.edges = append(toVertex.edges,from)
	}
	return g
}

func (g *Graph) Dfs(v *Vertex)int{
	myStack := &Stack{}
	myStack.Push(v.val)
	for myStack.Empty() == false {
		stackVertex := myStack.Pop()

	}
	return 0
}
