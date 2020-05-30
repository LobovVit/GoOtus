package lesson06

type List struct {
	len       int
	firstItem *Item
	lastItem  *Item
}

func (list *List) Len() int {
	return list.len
}

func (list *List) First() *Item {
	return list.firstItem
}

func (list *List) Last() *Item {
	return list.lastItem
}

func (list *List) PushFirst(v interface{}) {
	item := Item{nil, nil, v}
	if list.len == 0 {
		list.firstItem = &item
		list.lastItem = &item
	} else {
		item.next = list.firstItem
		list.firstItem.prev = &item
		list.firstItem = &item
	}
	list.len++
}

func (list *List) PushBack(v interface{}) {
	item := Item{nil, nil, v}
	if list.len == 0 {
		list.firstItem = &item
		list.lastItem = &item
	} else {
		item.prev = list.lastItem
		list.lastItem.next = &item
		list.lastItem = &item
	}
	list.len++
}

func (list *List) Remove(item Item) {
	next := item.next
	prev := item.prev
	if prev != nil {
		prev.next = next
	} else {
		list.firstItem = next
	}
	if next != nil {
		next.prev = prev
	} else {
		list.lastItem = prev
	}
	list.len--
	item.prev = nil
	item.next = nil
}
