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
	// Если есть первый элемент, то нужно изменить связи
	if lst.front != nil {
		lst.front.Next = itm
	}

	// Сверху добавляемый элемент и возврат
	lst.front = itm
	return itm
}

func (lst *list) pushBack(itm *ListItem) *ListItem {
	// нужно изменить связи
	lst.back.Prev = itm

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
	lst.length++
	return lst.pushFront(itm)
}

func (lst *list) PushBack(v interface{}) *ListItem {
	//Если нет первого элемента, можно использовать pushFront
	if lst.front == nil {
		return lst.PushFront(v)
	}
	lst.length++
	return lst.pushBack(&ListItem{v, lst.back, nil})
}

func NewList() List {
	return new(list)
}

func (lst *list) remove(i *ListItem) {
	// Если есть следующий элемент
	if i.Next != nil {
		// то связываем его с предыдущим (или nil)
		i.Next.Prev = i.Prev
	} else {
		// иначе связка списка с предыдущим
		lst.front = i.Prev
	}

	// Если есть предыдущий элемент
	if i.Prev != nil {
		// то связываем его со следющим
		i.Prev.Next = i.Next
	} else {
		// иначе связка списка со следующим (или nil)
		lst.back = i.Next
	}

	// Нужны ли строчки ниже? мне кажется нет
	// i.Prev = nil
	// i.Next = nil
}

func (lst *list) Remove(i *ListItem) {
	lst.remove(i)
	lst.length--
	// Нужна ли строчка ниже? мне кажется нет
	//i = nil
}

func (lst *list) MoveToFront(i *ListItem) {
	// Если впереди некого то это первый элемент
	if i.Next == nil {
		return
	}
	// Удалим элемент со старой позиции
	lst.remove(i)
	// Свяжем его с первым элементом(обратная связка в нутри pushFront)
	i.Prev = lst.front
	// Добавим элемент в верх
	lst.pushFront(i)
}
