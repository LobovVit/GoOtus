package lesson06

type Item struct {
	next *Item
	prev *Item
	data interface{}
}

func (item *Item) Value() interface{} {
	return item.data
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}
