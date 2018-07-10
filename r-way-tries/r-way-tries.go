package r_way_tries

import "log"

type Node struct {
	value interface{}
	next  map[string]*Node
}

var root = &Node{"", make(map[string]*Node)}
var size int

func Put(key string, val interface{}) {
	if key == "" {
		log.Fatal("Key not null");
	}
	put(root, key, val, 0)
}

func Get(key string) interface{} {
	if key == "" {
		log.Fatal("Key not null");
	}
	nodeResult := get(root, key, 0)
	return nodeResult.value
}

func Size() int  {
	return size
}

func get(node *Node, key string, d int) *Node {
	if d == len(key) {
		return node
	}
	c := string(key[d])
	return get(node.next[c], key, d+1)
}

func put(node *Node, key string, val interface{}, d int) *Node {
	if d == len(key) {
		if val == "" {
			size++
		}
		node.value = val
		return node
	}
	c := string(key[d])
	if _, ok := node.next[c]; !ok {
		node.next[c] = &Node{"", make(map[string]*Node)}
	}
	put(node.next[c], key, val, d+1)
	return node
}
