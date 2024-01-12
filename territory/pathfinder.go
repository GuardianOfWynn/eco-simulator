package territory

import (
	"fmt"
	"math"

	fibheap "github.com/starwander/gofibonacciheap"
	"github.com/victorbetoni/go-streams/streams"
)

type Pathfinder struct {
	From       *Territory
	Target     *Territory
	Claim      Claim
	RouteStyle RouteStyle
}

type Node struct {
	Territory *Territory
	Conns     []string
	Distance  float64
	Parent    *Node
}

func (n *Node) Tag() interface{} {
	return n.Territory
}

func (n *Node) Key() float64 {
	return n.Distance
}

func (p *Pathfinder) Route() []*Territory {

	//implement djikstra algorithm

	visited := map[string]bool{}
	nodes := map[string]*Territory{}
	distances := make(map[string]int)
	previous := make(map[string]string)
	root := p.From.Name
	nodesMap := map[string]*Node{}

	streams.StreamOf[*Territory](p.Claim.Territories...).ForEach(func(e *Territory) {
		nodesMap[e.Name] = &Node{
			Territory: e,
			Distance:  math.Inf(1),
			Conns:     e.Connections,
			Parent:    nil,
		}
		distances[e.Name] = math.MaxInt64
		nodes[e.Name] = e
	})

	nodesMap[root] = &Node{
		Territory: p.From,
		Conns:     p.From.Connections,
		Distance:  0,
		Parent:    nil,
	}

	distances[root] = 0
	heap := fibheap.NewFibHeap()
	heap.InsertValue(nodesMap[root])

	for heap.Num() != 0 {

		minimum := heap.ExtractMinValue()
		tag := fmt.Sprintf("%v", minimum.Tag())
		if visited[tag] {
			continue
		}

		visited[tag] = true
		fromNode := nodesMap[tag]

		for _, conn := range nodesMap[tag].Conns {
			if !visited[conn] {
				edgeWeight := 1 // if territory has tax, increase this value
				toNode := nodesMap[conn]
				if distances[fromNode.Territory.Name]+edgeWeight < distances[conn] {
					toNode.Distance = fromNode.Distance + float64(edgeWeight)
					toNode.Parent = fromNode
					distances[conn] = distances[fromNode.Territory.Name] + edgeWeight
					previous[conn] = fromNode.Territory.Name
					heap.InsertValue(toNode)
				}
			}
		}

	}

	path := []*Territory{}
	currentNode := previous[p.Target.Name]

	for currentNode != p.From.Name {
		path = append(path, nodes[currentNode])
		currentNode = previous[currentNode]
	}

	reversed := []*Territory{}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = path[j], path[i]
	}

	return reversed
}
