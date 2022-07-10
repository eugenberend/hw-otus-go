package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	FirstItem *ListItem
	LastItem  *ListItem
	length    int
}

func NewList() List {
	return new(list)
}

func (l *list) Front() *ListItem {
	return l.FirstItem
}

func (l *list) Back() *ListItem {
	return l.LastItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) PushFront(v interface{}) *ListItem {
	currentFront := l.Front()
	newFront := &ListItem{
		Value: v,
		Next:  currentFront,
		Prev:  nil,
	}
	l.FirstItem = newFront
	if l.Len() == 0 {
		l.LastItem = newFront
	} else {
		currentFront.Prev = newFront
	}
	l.length++
	return newFront
}

func (l *list) PushBack(v interface{}) *ListItem {
	currentBack := l.Back()
	newBack := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  currentBack,
	}
	l.LastItem = newBack
	if l.Len() == 0 {
		l.FirstItem = newBack
	} else {
		currentBack.Next = newBack
	}
	l.length++
	return newBack
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.FirstItem = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.LastItem = prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}
