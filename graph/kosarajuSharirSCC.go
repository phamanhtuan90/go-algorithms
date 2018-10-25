package graph

type KosarajuSharirSCC struct {
	marked [] bool
	count  int
	id     [] int
}

func NewKosarajuSharirSCC(G *DiGraph) *KosarajuSharirSCC {
	cc := KosarajuSharirSCC{
		make([] bool, G.V),
		0,
		make([] int, G.V),
	}
	depthFirstOrder := NewDepthFirstOrder(G.Reverse())
	posts := depthFirstOrder.ReversePost()
	var v int
	var err error
	for {
		v, err = posts.Pop()
		if err != nil {
			break
		}
		if !cc.marked[v] {
			cc.dfs(G, v)
			cc.count++
		}
	}

	return &cc
}

func (cc *KosarajuSharirSCC) dfs(G *DiGraph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for w, isConnect := range G.Adj(v) {
		if isConnect && !cc.marked[w] {
			cc.dfs(G, w)
		}
	}
}

func (cc *KosarajuSharirSCC) stronglyConnected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}
