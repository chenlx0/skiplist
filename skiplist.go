package skiplist

const (
	// DefaultLevel is default skip list level
	DefaultLevel = 8
)

// Node is a skip list node
type Node struct {
	element interface{}
	next    []*Node
}

// SkipList wrapper
type SkipList struct {
	maxLevel int
	head     *Node
	compare  *Comparable
}

// New return a new SkipList
func New(keyFunc *Comparable) *SkipList {
	return &SkipList{
		head:    nil,
		compare: keyFunc,
	}
}
