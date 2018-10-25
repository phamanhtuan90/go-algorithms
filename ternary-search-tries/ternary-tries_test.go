package ternary_search_tries


import (
	"testing"
	"fmt"
)

func TestPut(t *testing.T) {
	Put("sunday", 1);
	Put("monday", 2);
	Put("Pi", 3);
	Put("mondayy", 4);
	Put("mondayzzz", 5);

	fmt.Println(Get("mondayy"))
	fmt.Println(Get("Pi"))
}