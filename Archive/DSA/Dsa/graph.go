package dsa

import (
	"math"

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

func (g *Graph[T]) getAllVertices() []T {
	var vertices []T

	for key, _ := range g.AdjacencyList {
		vertices = append(vertices, key)
	}

	return vertices
}

// cycle detection
func (g *Graph[T]) CycleDetectionDfs(startNode T) bool {
	type color uint8

	const (
		WHITE color = iota
		GRAY
		BLACK
	)

	vertices := g.getAllVertices()
	vertexColor := make(map[T]color)

	for _, vertex := range vertices {
		vertexColor[vertex] = WHITE
	}

	var dfs func(T) bool
	var curr T
	var start *SllNode[T]
	stack := CreateStack[T](3)

	dfs = func(startNode T) bool {
		stack.Push(startNode)
		for stack.Length != 0 {
			curr = stack.Pop()
			vertexColor[curr] = GRAY
			start = g.AdjacencyList[curr].Head
			for start != nil {
				if vertexColor[start.Data] == WHITE {
					if dfs(start.Data) {
						return true
					}
				}

				start = start.Next
			}
		}

		return false
	}

	for key, value := range vertexColor {
		if value == WHITE {
			if dfs(key) {
				return true
			}
		}
	}

	return false
}

func (g *Graph[T]) CycleDetectionBfs() bool {
	vertices := g.getAllVertices()
	var curr T
	inDegree := make(map[T]uint)
	var start *SllNode[T]

	for _, vertex := range vertices {
		start = g.AdjacencyList[vertex].Head
		for start != nil {
			inDegree[start.Data] += 1
			start = start.Next
		}
	}

	q := CreateQueue[T](5)
	var res []T
	for key, value := range inDegree {
		if value == 0 {
			q.Append(key)
			res = append(res, key)
		}
	}

	for q.Length != 0 {
		curr = q.PopLeft()
		start = g.AdjacencyList[curr].Head
		for start != nil {
			inDegree[start.Data] -= 1
			if inDegree[start.Data] == 0 {
				q.Append(start.Data)
				res = append(res, start.Data)
			}
			start = start.Next
		}
	}

	return len(res) == len(vertices)
}

// topological sort
func (g *Graph[T]) TopologicalSort() []T {
	vertices := g.getAllVertices()
	var curr T
	inDegree := make(map[T]uint)
	var start *SllNode[T]

	for _, vertex := range vertices {
		start = g.AdjacencyList[vertex].Head
		for start != nil {
			inDegree[start.Data] += 1
			start = start.Next
		}
	}

	q := CreateQueue[T](5)
	var res []T
	for key, value := range inDegree {
		if value == 0 {
			q.Append(key)
			res = append(res, key)
		}
	}

	for q.Length != 0 {
		curr = q.PopLeft()
		start = g.AdjacencyList[curr].Head
		for start != nil {
			inDegree[start.Data] -= 1
			if inDegree[start.Data] == 0 {
				q.Append(start.Data)
				res = append(res, start.Data)
			}
			start = start.Next
		}
	}

	if len(res) != len(vertices) {
		panic("the graphic is cyclic")
	}

	return res
}

//single source shortest path

func (g *Graph[T]) DijkstraAlgo(startNode T) map[T]int {
	pq := CreatePriorityQueue[T]()
	var curr T
	var start *SllNode[T]
	pq.Insert(startNode)
	distances := make(map[T]int)
	vertices := g.getAllVertices()
	visited := make(map[T]struct{})
	for _, vertex := range vertices {
		distances[vertex] = math.MaxInt
	}

	distances[startNode] = 0

	for pq.Length != 0 {
		curr = pq.Dequeue()
		if _, exist := visited[curr]; !exist {
			start = g.AdjacencyList[curr].Head
			for start != nil {
				if distances[curr]+start.Weight < distances[start.Data] {
					distances[start.Data] = distances[curr] + start.Weight
				}

				start = start.Next
			}

			visited[curr] = struct{}{}
		}
	}

	return distances

}

func (g *Graph[T]) BellmanFord(startNode T) map[T]int {
	pq := CreateStack[T](5)
	var curr T
	var start *SllNode[T]
	pq.Push(startNode)
	distances := make(map[T]int)
	vertices := g.getAllVertices()
	visited := make(map[T]struct{})
	for _, vertex := range vertices {
		distances[vertex] = math.MaxInt
	}

	distances[startNode] = 0

	for pq.Length != 0 {
		curr = pq.Pop()
		if _, exist := visited[curr]; !exist {
			start = g.AdjacencyList[curr].Head
			for start != nil {
				if distances[curr]+start.Weight < distances[start.Data] {
					distances[start.Data] = distances[curr] + start.Weight
				}

				start = start.Next
			}

			visited[curr] = struct{}{}
		}
	}

	return distances
}

func (g *Graph[T]) FloydWarshall() {
	var adjacencyMatrix [][]int
	var start *SllNode[T]
	for i := range g.NumNodes {
		adjacencyMatrix = append(adjacencyMatrix, make([]int, g.NumNodes))
		for j := range g.NumNodes {
			adjacencyMatrix[i][j] = math.MaxInt
		}
	}
	vertices := g.getAllVertices()
	vertexIndex := make(map[T]int)
	for index, vertex := range vertices {
		vertexIndex[vertex] = index
	}

	for index, vertex := range vertices {
		start = g.AdjacencyList[vertex].Head
		for start != nil {
			adjacencyMatrix[index][vertexIndex[start.Data]] = start.Weight
			start = start.Next
		}
	}

	for i := range g.NumNodes {
		for j := range g.NumNodes {
			for k := range g.NumNodes {
				if adjacencyMatrix[i][j] < adjacencyMatrix[i][k]+adjacencyMatrix[k][j] {
					adjacencyMatrix[i][j] = adjacencyMatrix[i][k] + adjacencyMatrix[k][j]
				}
			}
		}
	}
}

// mst
func (g *Graph[T]) PrimsAlgo() {

}
