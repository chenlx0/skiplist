package skiplist

// Comparable defines a comparable element
type Comparable interface {
	Compare(lhs interface{}, rhs interface{}) bool
	Equals(lhs interface{}, rhs interface{}) bool
}
