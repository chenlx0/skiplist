package skiplist

type Comparable interface {
	Compare(lhs interface{}, rhs interface{}) bool
}
