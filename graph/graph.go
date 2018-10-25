package graph

type Graph struct {
	V   int
	Bag [][] bool
}

func NewGraph(v int) *Graph {
	G := Graph{
		v,
		make([][]bool, v, v),
	}
	for i := 0; i < v; i++ {
		G.Bag[i] = make([] bool, v)
	}

	return &G
}

func (G *Graph) AddEdge(v int, w int) {
	G.Bag[v][w] = true
	G.Bag[w][v] = true
}
func (G *Graph) Adj(v int) []bool  {
	return G.Bag[v]
}