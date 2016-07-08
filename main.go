package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MyHash struct {
	values *[]interface{}
	keys   *[]string
	elems  int
}

func New() *MyHash {
	vals := make([]interface{}, 5)
	keys := make([]string, 5)
	return &MyHash{
		&vals,
		&keys,
		0,
	}
}

func (h *MyHash) Size() int {
	return len(*h.values)
}

func (h *MyHash) Value(i uint32) interface{} {
	return (*h.values)[i]
}

func (h *MyHash) Key(i uint32) string {
	return (*h.keys)[i]
}

func (h *MyHash) Put(key string, val interface{}) {
	hash := hashKey(key)
	place := hashMod(hash, h.Size())

	if h.Value(place) != nil {
		h.expand()
		place = hash % uint32(h.Size())
	}

	(*h.keys)[place] = key
	(*h.values)[place] = val
	h.elems = h.elems + 1
}

func (h *MyHash) Get(key string) interface{} {
	hash := hashKey(key)
	place := hashMod(hash, h.Size())
	return h.Value(place)
}

func (h *MyHash) expand() {
	newSize := h.Size()*2 + 1
	oldValues := *h.values
	oldKeys := *h.keys
	newKeys := make([]string, newSize)
	newValues := make([]interface{}, newSize)

	for i, val := range oldValues {
		hash := hashKey(oldKeys[i])
		place := hashMod(hash, newSize)
		newKeys[place] = oldKeys[i]
		newValues[place] = val
	}

	h.keys = &newKeys
	h.values = &newValues

}

func (h *MyHash) Dump() {
	for i, _ := range *h.keys {
		j := uint32(i)
		if h.Value(j) != nil {
			val := fmt.Sprintf("%v", h.Value(j))
			fmt.Printf("%s: %s\n", h.Key(j), val)
		}
	}
}

func hashMod(hash uint32, size int) uint32 {
	return hash % uint32(size)
}

func hashKey(key string) uint32 {
	val := uint32(0)
	for _, b := range key {
		val += val*17 + uint32(b)
	}

	return val
}

func main() {
	rand.Seed(42)

	start := time.Now()
	for i := 0; i < 100; i++ {
		h := New()
		for {
			val := rand.Int()
			key := fmt.Sprintf("%v", val)
			h.Put(key, val)
			if h.elems > 1000 {
				break
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%dms\n", elapsed/1000000/100)
}
