1.package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("GoLang, Graph DFS and BFS implementation")
	fmt.Println("DFS : Depth First Search")
	fmt.Println("BFS : Breadth First Search")

	g := NewGraph()

	g.AddVertex("ajinkya")
	g.AddVertex("francesc")
	g.AddVertex("manish")
	g.AddVertex("albert")

	g.AddEdge("albert", "ajinkya")
	g.AddEdge("ajinkya", "albert")
	g.AddEdge("francesc", "ajinkya")
	g.AddEdge("francesc", "manish")
	g.AddEdge("manish", "francesc")
	g.AddEdge("manish", "albert")

	g.DFS("francesc")
	g.CreatePath("francesc", "albert")
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

type Graph struct {
	adjacency map[string][]string
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

func (g Graph) BFS(startingNode string) {
	visited := g.createVisited()
	var q []string

	visited[startingNode] = true
	q = append(q, startingNode)

	for len(q) > 0 {
		var current string
		current, q = q[0], q[1:]
		fmt.Println("BFS", current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}
}

func (g Graph) DFS(startingNode string) {
	visited := g.createVisited()
	g.dfsRecursive(startingNode, visited)
}

func (g Graph) dfsRecursive(startingNode string, visited map[string]bool) {
	visited[startingNode] = true
	fmt.Println("DFS", startingNode)
	for _, node := range g.adjacency[startingNode] {
		if !visited[node] {
			g.dfsRecursive(node, visited)
		}
	}
}

func (g Graph) CreatePath(firstNode, secondNode string) bool {
	visited := g.createVisited()
	var (
		path []string
		q    []string
	)
	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}
	fmt.Println("no link found")
	return false
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
GoLang, Graph DFS and BFS implementation
DFS : Depth First Search
BFS : Breadth First Search
DFS francesc
DFS ajinkya
DFS albert
DFS manish
francesc->ajinkya->albert

Program exited.
2.package main

import "fmt"

// FIFO is a FIFO queue
type FIFO struct {
	queue []interface{}
}

// New creates new FIFO and returns it
func New() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

// Push pushed node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := New()

	for _, val := range vals {
		fifo.Push(val)
	}

	for {
		front := fifo.Front()
		if front == nil {
			break
		}
		fmt.Println(front)
	}
}
1
2
3
4

Program exited.
3.package main

import (
	"container/list"
	"fmt"
)

func main() {
	vals := []int{1, 2, 3, 4,5,9,0,8}
	fifo := list.New()

	for _, val := range vals {
		fifo.PushBack(val)
	}

	for val := fifo.Front(); val != nil; val = val.Next() {
		fmt.Println(val.Value)
	}
}
1
2
3
4
5
9
0
8

Program exited.
4.package main

import (
	"fmt"
)

func main() {
	fmt.Println("GoLang, Graph DFS and BFS implementation")
	fmt.Println("DFS : Depth First Search")
	fmt.Println("BFS : Breadth First Search")

	g := NewGraph()

	g.AddVertex("ajinkya")
	g.AddVertex("francesc")
	g.AddVertex("manish")
	g.AddVertex("albert")

	g.AddEdge("albert", "ajinkya")
	g.AddEdge("ajinkya", "albert")
	g.AddEdge("francesc", "ajinkya")
	g.AddEdge("francesc", "manish")
	g.AddEdge("manish", "francesc")
	g.AddEdge("manish", "albert")

	g.BFS("francesc")
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}
type Graph struct {
	adjacency map[string][]string
}


func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

func (g Graph) BFS(startingNode string) {
	visited := g.createVisited()
	var q []string

	visited[startingNode] = true
	q = append(q, startingNode)

	for len(q) > 0 {
		var current string
		current, q = q[0], q[1:]
		fmt.Println("visiting node", current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}


func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
} 
GoLang, Graph DFS and BFS implementation
DFS : Depth First Search
BFS : Breadth First Search
visiting node francesc
visiting node ajinkya
visiting node manish
visiting node albert

Program exited.