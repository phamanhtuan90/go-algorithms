package graph

type DiGraph struct {
	V   int
	Bag [][] bool
}

func NewDiGraph(v int) *DiGraph {
	G := DiGraph{
		v,
		make([][]bool, v, v),
	}
	for i := 0; i < v; i++ {
		G.Bag[i] = make([] bool, v)
	}

	return &G
}

func (G *DiGraph) AddEdge(v int, w int) {
	G.Bag[v][w] = true
}
func (G *DiGraph) Adj(v int) []bool  {
	return G.Bag[v]
}

func (G *DiGraph) Reverse() *DiGraph  {
	reverse := NewDiGraph(G.V)
	for v := 0; v < G.V; v++ {
		for w, isConnect := range G.Adj(v) {
			if isConnect {
				reverse.AddEdge(w, v)
			}
		}
	}
	return reverse

}



