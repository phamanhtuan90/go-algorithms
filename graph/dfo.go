package graph

import "github.com/go-algorithms/stack"

type DepthFirstOrder struct {
	marked      [] bool
	reversePost *stack.Stack
}

func NewDepthFirstOrder(G DiGraph) *DepthFirstOrder {
	d := DepthFirstOrder{
		make([]bool, G.V),
		stack.NewStack(),
	}
	for v := 0; v < G.V; v++ {
		if !d.marked[v] {
			d.dfs(G, v)
		}
	}
	return &d
}

func (d *DepthFirstOrder) dfs(G DiGraph, v int) {
	d.marked[v] = true
	for w, value := range G.Adj(v) {
		if value && !d.marked[w] {
			d.dfs(G, w)

		}
	}
	d.reversePost.Push(v)
}

func (d *DepthFirstOrder) ReversePost() *stack.Stack {
	return d.reversePost
}
