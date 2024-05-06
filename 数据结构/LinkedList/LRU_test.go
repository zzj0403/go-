package _linkedList

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := Constructor(3)
	t.Log(cache)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2))
	cache.PrintCache()
	cache.Put(3, 3)
	cache.PrintCache()
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(1))
	cache.Put(4, 4)
	cache.Put(5, 5)
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(1))
	cache.PrintCache()

}
