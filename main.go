package main

import (
	"log"
	"strings"
	"fmt"
)

type Node struct {
	c             string
	value         interface{}
	mid, lft, rgt *Node
}

var root = &Node{"", "", &Node{}, &Node{}, &Node{}}
var size int

func Put(key string, val interface{}) {
	if key == "" {
		log.Fatal("Key not null")
	}
	//if !contains(key){
	//	size++
	//}
	put(root, key, val, 0)

}

func Get(key string) interface{} {
	if key == "" {
		log.Fatal("Key not null")
	}
	nodeResult := get(root, key, 0)
	if nodeResult == nil {
		return nil
	}
	return nodeResult.value
}

func Size() int {
	return size
}

func contains(key string) bool {
	if key == "" {
		log.Fatal("Key not null")
	}
	return get(root, key, 0) != nil
}

func put(node *Node, key string, val interface{}, d int) *Node {
	c := string(key[d])
	//if node == nil {
	//	node = &Node{c, "", &Node{}, &Node{}, &Node{}}
	//}

	if node.c == "" {
		node.c = c
		node.lft, node.rgt, node.mid = &Node{}, &Node{}, &Node{}
	}
	//fmt.Println( strings.Compare(c, node.c), c)
	if strings.Compare(c, node.c) == -1 {
		put(node.lft, key, val, d)
	} else if strings.Compare(c, node.c) == 1 {
		put(node.rgt, key, val, d)
	} else if d < len(key)-1 {
		put(node.mid, key, val, d+1)
	} else {
		node.value = val
	}
	return node
}

func get(node *Node, key string, d int) *Node {
	if node == nil {
		return nil
	}
	c := string(key[d])
	if c < node.c {
		return get(node.lft, key, d)
	} else if c > node.c {
		return get(node.rgt, key, d)
	} else if d < len(key)-1 {
		return get(node.mid, key, d+1)
	} else {
		return node
	}

}

func main() {
	Put("sunday", 1);
	Put("monday", 2);
	Put("Pi", 3);
	Put("mondayy", 4);
	Put("mondayzzz", 5);
	fmt.Println(Get("mondayy"))
	fmt.Println(Get("Pi"))
}
