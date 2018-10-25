package graph

import (
	"github.com/Workiva/go-datastructures/queue"
	"math"
)

type BFS struct {
	marked [] bool
	egdeTo [] int
	distTo [] int
	s      int
}

func NewBFS(G Graph, s int) *BFS {
	b := BFS{
		make([] bool, G.V),
		make([] int, G.V),
		make([] int, G.V),
		s,
	}
	return &b
}

func (B *BFS) Bfs(G Graph, s int) {
	for i := 0; i < G.V; i++ {
		B.distTo[i] = math.MaxInt32
	}
	B.distTo[s] = 0
	B.marked[s] = true
	q := queue.New(int64(G.V))
	q.Put(s)
	for !q.Empty() {
		v, _ := q.Peek()
		for w, isConnect := range G.Adj(v.(int)) {
			if isConnect && !B.marked[w] {
				B.marked[w] = true
				B.distTo[w] = B.distTo[v.(int)] + 1
				B.egdeTo[w] = v.(int)
				q.Put(w)
			}
		}
	}
}
