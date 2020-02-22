package skiplist

import (
	"testing"
)

type Integer struct{}

func (i *Integer) Compare(lhs interface{}, rhs interface{}) bool {
	a := lhs.(int32)
	b := rhs.(int32)
	return a > b
}

func (i *Integer) Equals(lhs interface{}, rhs interface{}) bool {
	return lhs == rhs
}

func TestAdd(t *testing.T) {
	list := New(new(Integer))

	for i := 50; i > 0; i-- {
		num := int32(i)
		list.Add(num)
	}

	var x int32 = 50
	t.Log(list.Contain(x))

	t.Log("size of list: ", list.Len())

	for i := 0; i < 50; i++ {
		a := list.Next()
		t.Log(a.(int32))
	}
}
