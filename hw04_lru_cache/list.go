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
	length int
	front  *ListItem
	back   *ListItem
}

func (lst *list) Len() int {
	return lst.length
}

func (lst *list) Front() *ListItem {
	return lst.front
}

func (lst *list) Back() *ListItem {
	return lst.back
}

func (lst *list) pushFront(itm *ListItem) *ListItem {
	// Если будет проблема со списком
	lst.length++

	// Если есть первый элемент, то нужно изменить связи
	if lst.front != nil {
		lst.front.Next = itm
		//itm.Prev = lst.front
	}

	// Сверху добавляемый элемент и возврат
	lst.front = itm
	return itm
}

func (lst *list) pushBack(itm *ListItem) *ListItem {
	// Если будет проблема со списком
	lst.length++

	// нужно изменить связи
	lst.back.Prev = itm
	//itm.Next = lst.back

	// Добавляем элемент и возврат
	lst.back = itm
	return itm
}

func (lst *list) PushFront(v interface{}) *ListItem {
	itm := &ListItem{v, nil, lst.front}
	//Если нет последнего элемента, то это последний
	if lst.back == nil {
		lst.back = itm
	}
	return lst.pushFront(itm)
}

func (lst *list) PushBack(v interface{}) *ListItem {
	//Если нет первого элемента, можно использовать pushFront
	if lst.front == nil {
		return lst.PushFront(v)
	}

	return lst.pushBack(&ListItem{v, lst.back, nil})
}

func NewList() List {
	return new(list)
}

func (lst *list) remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		lst.front = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		lst.back = i.Next
	}

	i.Prev = nil
	i.Next = nil
}

func (lst *list) Remove(i *ListItem) {
	lst.remove(i)
	lst.length--
	i = nil
}

func (lst *list) MoveToFront(i *ListItem) {
	if i.Next == nil {
		return
	}
	lst.remove(i)
	i.Prev = lst.front
	lst.pushFront(i)
}
