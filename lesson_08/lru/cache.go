package lru

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type Cache struct {
	capacity int
	Head     *Node
	Tail     *Node
	dataMap  map[string]*Node
}

type Node struct {
	Key   string
	Value string
	Next  *Node
	Prev  *Node
}

func NewLruCache(capacity int) LruCache {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	return &Cache{
		capacity: capacity,
		Head:     nil,
		Tail:     nil,
		dataMap:  make(map[string]*Node, capacity),
	}
}

func (c *Cache) Put(key, value string) {
	currentCacheCapacity := len(c.dataMap)

	existingNodeByKey, ok := c.dataMap[key]

	currentHead := c.Head

	if ok {
		c.Head = existingNodeByKey

		tempNext := existingNodeByKey.Next
		tempPrev := existingNodeByKey.Prev

		if tempPrev != nil {
			tempPrev.Next = tempNext
		}

		if tempNext != nil {
			tempNext.Prev = tempPrev
		}

		existingNodeByKey.Value = value
		existingNodeByKey.Next = currentHead
		existingNodeByKey.Prev = nil

		if existingNodeByKey == c.Tail {
			c.Tail = tempPrev
		}
	} else {
		if currentCacheCapacity+1 > c.capacity {
			if c.Tail == nil {
				panic("Tail is empty")
			}

			tailNode, ok := c.dataMap[c.Tail.Key]

			if !ok {
				panic("Tail node not found")
			}

			tempPrev := tailNode.Prev

			if tempPrev != nil {
				tempPrev.Next = nil
				c.Tail = tempPrev
			}

			delete(c.dataMap, tailNode.Key)
		}

		newNode := &Node{
			Key:   key,
			Value: value,
			Next:  currentHead,
			Prev:  nil,
		}
		c.dataMap[key] = newNode
		c.Head = newNode

		if currentHead != nil {
			currentHead.Prev = newNode
		}

		if c.Tail == nil {
			c.Tail = newNode
		}
	}
}

func (c *Cache) Get(key string) (string, bool) {
	existingNodeByKey, ok := c.dataMap[key]

	if !ok {
		return "", false
	}

	if c.Head != existingNodeByKey {
		currentHead := c.Head
		c.Head = existingNodeByKey

		tempNext := existingNodeByKey.Next
		tempPrev := existingNodeByKey.Prev

		if tempPrev != nil {
			tempPrev.Next = tempNext
		}

		if tempNext != nil {
			tempNext.Prev = tempPrev
		}

		existingNodeByKey.Next = currentHead
		existingNodeByKey.Prev = nil

		if existingNodeByKey == c.Tail {
			c.Tail = tempPrev
		}
	}

	return existingNodeByKey.Value, true
}
