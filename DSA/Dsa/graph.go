package dsa

import (
	"golang.org/x/exp/constraints"
)

type Graph[T constraints.Ordered] struct {
	AdjacencyList map[T]*Sll[T]
	NumNodes      uint
	directed      bool
}

func (g *Graph[T]) CreateGraph() *Graph[T] {
	return &Graph[T]{AdjacencyList: make(map[T]*Sll[T]), NumNodes: 0}
}

func (g *Graph[T]) AddVertex(data T) {
	g.AdjacencyList[data] = &Sll[T]{}
	g.NumNodes += 1
}

func (g *Graph[T]) AddEdge(from T, to T) {
	_, fromExists := g.AdjacencyList[from]
	_, toExists := g.AdjacencyList[to]
	if fromExists {
		g.AdjacencyList[from] = &Sll[T]{}
		g.NumNodes += 1
	}

	if toExists {
		g.AdjacencyList[to] = &Sll[T]{}
		g.NumNodes += 1
	}

	g.AdjacencyList[from].Append(to)
	if !g.directed {
		g.AdjacencyList[to].Append(from)
	}
}

func (g *Graph[T]) DfsIterative(startNode T) []T {
	if _, startNodeExist := g.AdjacencyList[startNode]; !startNodeExist {
		panic("startNode does not exist")
	}

	visited := make(map[T]struct{})
	stack := Stack[T]{}
	var res []T

	stack.Push(startNode)
	var curr T

	for stack.Length != 0 {
		curr = stack.Pop()
		visited[curr] = struct{}{}
		res = append(res, curr)
		start := g.AdjacencyList[curr].Head
		for start != nil {
			stack.Push(start.Data)
			start = start.Next
		}
	}

	for key := range g.AdjacencyList {
		if _, exist := visited[key]; !exist {
			func() {
				stack.Push(key)
				for stack.Length != 0 {
					curr = stack.Pop()
					visited[curr] = struct{}{}
					res = append(res, curr)
					start := g.AdjacencyList[curr].Head
					for start != nil {
						stack.Push(start.Data)
						start = start.Next
					}
				}
			}()
		}
	}

	return res
}

func (g *Graph[T]) BfsIterative(startNode T) []T {
	if _, startNodeExist := g.AdjacencyList[startNode]; !startNodeExist {
		panic("startNode does not exist")
	}

	visited := make(map[T]struct{})
	queue := Queue[T]{}
	var res []T

	queue.Append(startNode)
	var curr T

	for queue.Length != 0 {
		curr = queue.PopLeft()
		visited[curr] = struct{}{}
		res = append(res, curr)
		start := g.AdjacencyList[curr].Head
		for start != nil {
			queue.Append(start.Data)
			start = start.Next
		}
	}

	for key := range g.AdjacencyList {
		if _, exist := visited[key]; !exist {
			func() {
				queue.Append(key)
				for queue.Length != 0 {
					curr = queue.PopLeft()
					visited[curr] = struct{}{}
					res = append(res, curr)
					start := g.AdjacencyList[curr].Head
					for start != nil {
						queue.Append(start.Data)
						start = start.Next
					}
				}
			}()
		}
	}

	return res
}

func (g *Graph[T]) NumIslands(startNode T) uint {
	if _, startNodeExist := g.AdjacencyList[startNode]; !startNodeExist {
		panic("startNode does not exist")
	}

	visited := make(map[T]struct{})
	stack := Stack[T]{}

	stack.Push(startNode)
	var curr T

	for stack.Length != 0 {
		curr = stack.Pop()
		visited[curr] = struct{}{}
		start := g.AdjacencyList[curr].Head
		for start != nil {
			stack.Push(start.Data)
			start = start.Next
		}
	}

	var numIslands uint = 1

	for key := range g.AdjacencyList {
		if _, exist := visited[key]; !exist {
			numIslands += 1
			func() {
				stack.Push(key)
				for stack.Length != 0 {
					curr = stack.Pop()
					visited[curr] = struct{}{}
					start := g.AdjacencyList[curr].Head
					for start != nil {
						stack.Push(start.Data)
						start = start.Next
					}
				}
			}()
		}
	}

	return numIslands
}

//cycle detection
//topological sort
//single source shortest path
//mst
