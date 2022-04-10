package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	// Попробуем получить элемент
	val, ok := cache.items[key]

	if !ok {
		return nil, ok
	}

	// Если если он есть, то поднимем его вверх
	cache.queue.MoveToFront(val)
	return val, ok
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	// Если элемент есть то обновим ему значение
	if val, ok := cache.Get(key); ok {
		// Я не смог найти, как можно было бы добратся до value
		// и изменить его не пересоздавая объект, это возможно?
		val.(*ListItem).Value = cacheItem{key, value}
		return ok
	}

	// Если кэш заполнен, удалим последний элемент
	if cache.capacity == cache.queue.Len() {
		delete(cache.items, cache.queue.Back().Value.(cacheItem).key)
		cache.queue.Remove(cache.queue.Back())
	}

	// Добавим элемент в самый верх
	cache.items[key] = cache.queue.PushFront(cacheItem{key, value})

	return false
}

func (cache *lruCache) Clear() {
	// Можно в место цикла просто присвоить хешу новую hesh_map и создать новую очередь?
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}
