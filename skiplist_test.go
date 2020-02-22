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

func TestAdd1(t *testing.T) {
	list := New(new(Integer))

	for i := 500; i >= 0; i -= 5 {
		t.Log(i)
		num := int32(i)
		list.Add(num)
	}

	if list.Len() != 101 {
		t.Error("list size must be 101, but ", list.Len(), " get")
	}

	if !list.Contain(int32(500)) {
		t.Error("`500` not in list")
	}

	if !list.Contain(int32(500 - 5*55)) {
		t.Error("`500-5*55` not in list")
	}

	if !list.Contain(int32(0)) {
		t.Error("`0` not in list")
	}
}

func TestAddAndRemove(t *testing.T) {
	list := New(new(Integer))

	list.Add(int32(60))
	list.Add(int32(70))
	list.Add(int32(80))
	list.Add(int32(80))
	list.Add(int32(80))
	list.Add(int32(81))
	list.Add(int32(2))

	if !list.Contain(int32(2)) {
		t.Error("`2` not in list")
	}

	if !list.Contain(int32(80)) {
		t.Error("`80` not in list")
	}

	list.Remove(int32(2))
	list.Remove(int32(80))

	if list.Contain(int32(2)) {
		t.Error("must not contain `2`")
	}

	if list.Contain(int32(80)) {
		t.Error("must not contain `80`")
	}
}
