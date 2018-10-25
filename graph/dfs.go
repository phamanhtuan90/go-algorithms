package graph

import stack2 "github.com/go-algorithms/stack"

type DFS struct {
	marked [] bool
	egdeTo [] int
	s      int
}

func NewDFS(G Graph, s int) *DFS {
	d := DFS{
		make([]bool, G.V),
		make([]int, G.V),
		s,
	}
	d.Dfs(G, s)
	return &d
}

func (d *DFS) Dfs(G Graph, v int) {
	d.marked[v] = true
	for w, value := range G.Adj(v) {
		if value && !d.marked[w] {
			d.Dfs(G, w)
			d.egdeTo[w] = v
		}
	}
}

func (d *DFS) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DFS) PathTo(v int) interface{} {
	stack := stack2.NewStack()
	if d.HasPathTo(v) {
		for x := v; x != d.s; x = d.egdeTo[x] {
			stack.Push(x)
		}
		stack.Push(d.s)
	}

	return stack

}
