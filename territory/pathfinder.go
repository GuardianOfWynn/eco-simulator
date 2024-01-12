package territory

import (
	"fmt"
	"math"
	"time"

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
	Territory string
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

// for cheapest: order conns by tax value
func (p *Pathfinder) Route() []string {
	fmt.Println("começou")
	//implement djikstra algorithm
	start := time.Now().UnixMilli()
	root := p.From.Name
	nodesMap := map[string]*Node{}
	nodesMap[root] = &Node{
		Territory: root,
		Conns:     p.From.Connections,
		Distance:  0, Parent: nil,
	}
	streams.StreamOf[*Territory](p.Claim.Territories...).ForEach(func(e *Territory) {
		nodesMap[e.Name] = &Node{
			Territory: e.Name,
			Distance:  math.Inf(1),
			Conns:     e.Connections,
			Parent:    nil,
		}
	})

	path := []string{}
	heap := fibheap.NewFibHeap()

	for _, v := range nodesMap {
		heap.InsertValue(v)
	}

	for heap.Num() != 0 {
		minimum := heap.ExtractMinValue()
		tag := fmt.Sprintf("%v", minimum.Tag())
		path = append(path, tag)
		for _, conn := range nodesMap[tag].Conns {
			p.relax(nodesMap[tag], nodesMap[conn], 1)
		}
	}
	end := time.Now().UnixMilli()
	fmt.Println("took ", start-end, " milliseconds")

	return path
}

func (p *Pathfinder) relax(from, to *Node, edgeWeight int) {
	if to == nil {
		fmt.Println(from.Territory)
	}
	if from.Distance+float64(edgeWeight) < to.Distance {
		to.Distance = from.Distance + float64(edgeWeight)
		to.Parent = from
	}
}
