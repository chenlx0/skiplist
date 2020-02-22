package skiplist

import (
	"math/rand"
)

const (
	// DefaultMaxLevel is default skip list level
	DefaultMaxLevel = 18
)

// Node is a skip list node
type Node struct {
	val  interface{}
	next []*Node
}

// SkipList wrapper
type SkipList struct {
	maxLevel      int
	size          int
	preNodesCache []*Node
	heads         []*Node
	nextNode      *Node
	compare       Comparable
}

// New returns a new SkipList
func New(keyFunc Comparable) *SkipList {
	return &SkipList{
		maxLevel:      DefaultMaxLevel,
		size:          0,
		preNodesCache: make([]*Node, DefaultMaxLevel),
		heads:         make([]*Node, DefaultMaxLevel),
		compare:       keyFunc,
	}
}

// NewWithLevel also returns a new SkipList
// But you can specify its max level
func NewWithLevel(keyFunc Comparable, level int) *SkipList {
	return &SkipList{
		maxLevel:      level,
		size:          0,
		preNodesCache: make([]*Node, level),
		heads:         make([]*Node, level),
		compare:       keyFunc,
	}
}

// randLevel generates random node level
func (list *SkipList) randLevel() int {
	randFloat := rand.Float32()

	var threshold float32 = 1.0
	for i := 1; i < list.maxLevel; i++ {
		threshold = threshold / 2
		if randFloat > threshold {
			return i
		}
	}

	return list.maxLevel
}

func (list *SkipList) getPrevNodes(val interface{}) []*Node {
	prevs := list.preNodesCache
	var next *Node

	for i := 0; i < len(prevs); i++ {
		prevs[i] = nil
	}

	for i := list.maxLevel - 1; i >= 0; i-- {
		if prevs[i] == nil {
			next = list.heads[i]
		} else {
			next = prevs[i]
		}

		for next != nil && list.compare.Compare(val, next.val) {
			prevs[i] = next
			next = next.next[i]
		}

		if i != 0 && prevs[i] != nil {
			prevs[i-1] = prevs[i]
		}
	}

	return prevs
}

// Next returns value in list iterally
func (list *SkipList) Next() interface{} {
	if list.heads[0] == nil {
		return nil
	}

	if list.nextNode == nil {
		list.nextNode = list.heads[0]
	}

	ret := list.nextNode.val
	list.nextNode = list.nextNode.next[0]
	return ret
}

// Len returns list size
func (list *SkipList) Len() int {
	return list.size
}

// Contain returns if list contains a such value
func (list *SkipList) Contain(val interface{}) bool {
	if list.size == 0 {
		return false
	}

	if list.compare.Equals(list.heads[0].val, val) {
		return true
	}

	prevs := list.getPrevNodes(val)

	if prevs[0] == nil || prevs[0].next[0] == nil {
		return false
	}

	return list.compare.Equals(val, prevs[0].next[0].val)
}

// Add adds a value to list
func (list *SkipList) Add(val interface{}) {
	level := list.randLevel()
	newNode := &Node{
		val:  val,
		next: make([]*Node, level),
	}

	// search all preNodes
	prevs := list.getPrevNodes(val)
	for i := 0; i < level; i++ {
		if prevs[i] == nil {
			newNode.next[i] = list.heads[i]
			list.heads[i] = newNode
		} else {
			newNode.next[i], prevs[i].next[i] = prevs[i].next[i], newNode
		}
	}

	list.size++
}

// Remove removes all elements equal to the value
func (list *SkipList) Remove(val interface{}) {
	prevs := list.getPrevNodes(val)
	var start *Node

	for i := 0; i < len(prevs); i++ {
		start = prevs[i]
		if start == nil || start.next[i] == nil ||
			!list.compare.Equals(start.next[i].val, val) {
			continue
		}

		start = start.next[i]
		for list.compare.Equals(start.val, val) {
			start = start.next[i]
		}

		prevs[i].next[i] = start
	}

	// remove nodes in head
	for i := 0; i < len(list.heads); i++ {
		if list.heads[i] != nil &&
			list.compare.Equals(list.heads[i].val, val) {
			list.heads[i] = list.heads[i].next[i]
		}
	}

	list.size--
}
