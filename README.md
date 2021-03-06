# SkipList

This is a skip list implementation in Go. (Sorted List)

Go 语言的跳跃表实现（有序集合）

To start using this library:

```
go get github.com/chenlx0/skiplist
```



Example:

```Go
// Define compare functions before you create a skip list instance
type Integer struct{}

func (i *Integer) Compare(lhs interface{}, rhs interface{}) bool {
	return lhs.(int32) > rhs.(int32)
}

func (i *Integer) Equals(lhs interface{}, rhs interface{}) bool {
	return lhs == rhs
}

list := skiplist.New(new(Integer))

// Add an element
list.Add(int32(100))
list.Add(int32(200))
list.Add(int32(300))

// Get if list contains such an element
if list.Contain(int32(100)) {
    // Remove element
    fmt.Println("ok")
}

// Traverse all elements, and they are sorted!
for i := 0; i < list.Len(); i++ {
    fmt.Printf("Get %d", list.Next())
}

```