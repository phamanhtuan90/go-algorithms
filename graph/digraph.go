package graph

type DiGraph struct {
	V   int
	Bag [][] bool
}

func New(v int) *DiGraph {
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
	G.Bag[w][v] = true
}
func (G *DiGraph) Adj(v int) []bool  {
	return G.Bag[v]
}