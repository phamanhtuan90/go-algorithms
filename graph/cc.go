package graph

type CC struct {
	marked [] bool
	count  int
	id     [] int
}

func NewCC(G Diraph) *CC {
	cc := CC{
		make([] bool, G.V),
		0,
		make([] int, G.V),
	}
	for i := 0; i < G.V; i++ {
		if !cc.marked[i] {
			cc.dfs(G, i)
			cc.count++
		}
	}
	return &cc
}

func (cc *CC) dfs(G Diraph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for w, isConnect := range G.Adj(v){
		if isConnect && !cc.marked[w] {
			cc.dfs(G, w)
		}
	}
}
