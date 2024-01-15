package territory

import (
	"fmt"
	"math"

	fibheap "github.com/starwander/gofibonacciheap"
	"github.com/victorbetoni/go-streams/streams"
)

type Pathfinder struct {
	Root     *Territory
	GuildMap *GuildMap
}

type Node struct {
	Territory *Territory
	Conns     []string
	Distance  float64
	Parent    *Node
}

func (n *Node) Tag() interface{} {
	return n.Territory.Name
}

func (n *Node) Key() float64 {
	return n.Distance
}

func (p *Pathfinder) djikstra(target *Territory, style RouteStyle) (map[string]int, map[string]string, map[string]*Territory) {

	visited := map[string]bool{}
	nodes := map[string]*Territory{}
	distances := make(map[string]int)
	previous := make(map[string]string)
	root := p.Root.Name
	nodesMap := map[string]*Node{}

	streams.StreamOf[*Territory](p.GuildMap.Territories...).ForEach(func(e *Territory) {
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
		Territory: p.Root,
		Conns:     p.Root.Connections,
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
		for _, conn := range nodes[tag].Connections {
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

	return distances, previous, nodes
}

func (p *Pathfinder) Route(target *Territory, style RouteStyle) []*Territory {

	//djikstra algorithm

	_, previous, nodes := p.djikstra(target, style)

	path := []*Territory{}
	currentNode := previous[target.Name]

	for currentNode != p.Root.Name {
		path = append(path, nodes[currentNode])
		currentNode = previous[currentNode]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (p *Pathfinder) GetDistance(target *Territory) int64 {
	distance, _, _ := p.djikstra(target, FASTEST)
	return int64(distance[target.Name])
}
