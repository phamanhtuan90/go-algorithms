package main

import (
	"github.com/go-algorithms/graph"
	"fmt"
)

func main()  {
	G := graph.New(4)
	G.AddEdge(1 , 2)
	fmt.Println(G.Adj(1))
}
