package hashes

// Implementation of simple chaining hash.
// Hash Key is a string and hash Value is limited to integer.
// Hash has fixed size and does not support rehashing.

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const SIZE uint32 = 100

type Node struct {
	Next *Node
	Key  string
	Val  int
}

type SCHashTable struct {
	data []*Node
}

func NewSCHashTable() SCHashTable {
	newTable := SCHashTable{}
	newTable.data = make([]*Node, SIZE)

	return newTable
}

func (h *SCHashTable) Set(key string, val int) {
	index := getIndex(key)
	node := h.Find(key, h.data[index])
	if node == nil {
		h.data[index] = &Node{h.data[index], key, val}
	} else {
		node.Val = val
	}
}

func (h *SCHashTable) Get(key string) (int, error) {
	temp := h.Find(key, h.data[getIndex(key)])

	if temp != nil {
		return temp.Val, nil
	} else {
		return 0, errors.New("No value")
	}
}

func (h *SCHashTable) Del(key string) {
	index := getIndex(key)
	node := h.data[index]
	var prev *Node = nil

	for node != nil {
		if node.Key == key {
			if prev == nil {
				h.data[index] = node.Next
				fmt.Println(node.Next)
				fmt.Println(h.data[index])
			} else {
				prev.Next = node.Next
			}
			return
		}
		prev, node = node, node.Next
	}
}

func (h *SCHashTable) Find(key string, node *Node) *Node {
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil
}

func getIndex(key string) uint32 {
	h := fnv.New32()
	h.Write([]byte(key))
	return h.Sum32() % SIZE
}
