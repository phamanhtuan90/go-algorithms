package graph

import "github.com/go-algorithms/stack"

type DirectDFS struct {
	marked [] bool
	egdeTo [] int
	s      int
}

func NewDirectDFS(G *DiGraph, s int) *DirectDFS {
	d := DirectDFS{
		make([]bool, G.V),
		make([]int, G.V),
		s,
	}
	d.dfs(G, s)
	return &d
}

func (d *DirectDFS) dfs(G *DiGraph, v int) {
	d.marked[v] = true
	for w, value := range G.Adj(v) {
		if value && !d.marked[w] {
			d.dfs(G, w)
			d.egdeTo[w] = v
		}
	}
}

func (d *DirectDFS) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DirectDFS) PathTo(v int) interface{} {
	paths := stack.NewStack()
	if d.HasPathTo(v) {
		for x := v; x != d.s; x = d.egdeTo[x] {
			paths.Push(x)
		}
		paths.Push(d.s)
	}

	return paths

}
