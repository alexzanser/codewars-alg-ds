package main

import "fmt"

type Graph struct {
	ID		string
	kids	map[string]Graph
	weight	int
}


func findLowestCost(g Graph, costs map[string]int) (minCostNode Graph) {
	min := 1000000000
	minCostNode = Graph{}
	for _, node := range g.kids {
		if costs[node.ID] < min && costs[node.ID] > 0 {
			min = costs[node.ID]
			minCostNode = node
		}
	}
	return minCostNode
}

func main() {
	start := Graph {ID :"start", weight: 0}
	start.kids = map[string]Graph{
		"a":  Graph {ID:"a", weight: 6},
		"b" : Graph {ID:"b", weight: 2},
	}
	if a , ok := start.kids["a"]; ok {
		a.kids = map[string]Graph {
			"fin" : Graph{ID:"fin", weight: 1},
		}
	}
	if b, ok := start.kids["b"]; ok {
		b.kids = map[string]Graph {
			"a" : Graph{ID:"a", weight: 3},
			"fin" : Graph{ID:"fin", weight: 5},
		}
	}

	costs := map[string]int {
		"a" : 6,
		"b" : 2,
		"finn" : -1,
	}

	parents := map[string]string {
		"a" : "start",
		"b" : "start",
		"in" : "",
	}
	processed := make([]Graph, 0)

	node := findLowestCost(start, costs)

	for node.ID != "" {
		cost := costs[node.ID]
		for _, n := range node.kids {
			newCost := cost + n.weight
			if n.weight > newCost {
				n.weight = newCost
			}
			parents[n.ID] = node.ID
		}
		processed = append(processed, node)
	}
}