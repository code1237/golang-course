package lru

import (
	"time"
)

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type Cache struct {
	capacity int
	list     List
}

type List struct {
	Head *ListItem
}

func (l *Cache) Put(key, value string) {
	if l.list.Head == nil {
		l.list.Head = &ListItem{
			Key:    key,
			Value:  value,
			Next:   nil,
			Prev:   nil,
			ReadAt: time.Now(),
		}

		return
	}

	listLength := 0

	current := l.list.Head

	for i := 0; i < l.capacity; i++ {
		listLength++

		if current.Key == key {
			if listLength == 1 {
				l.list.Head.ReadAt = time.Now()
				l.list.Head.Value = value
				return
			}

			currentNext := current.Next
			tempHead := l.list.Head
			l.list.Head = &ListItem{
				Key:    key,
				Value:  value,
				Next:   tempHead,
				Prev:   nil,
				ReadAt: time.Now(),
			}

			tempHead.Next = currentNext
			tempHead.Prev = l.list.Head

			return
		}

		if current.Next != nil {
			current = current.Next
			continue
		}

		break
	}

	tempHead := l.list.Head
	l.list.Head = &ListItem{
		Key:    key,
		Value:  value,
		Next:   tempHead,
		Prev:   nil,
		ReadAt: time.Now(),
	}

	tempHead.Prev = l.list.Head

	if listLength >= l.capacity {
		current.Prev.Next = nil
	}
}

func (l *Cache) Get(key string) (string, bool) {
	current := l.list.Head

	for i := 0; i < l.capacity; i++ {
		if current.Key == key {
			l.Put(key, current.Value)
			return current.Value, true
		}

		current = current.Next
	}

	return "", false
}

type ListItem struct {
	Key    string
	Value  string
	Next   *ListItem
	Prev   *ListItem
	ReadAt time.Time
}

func NewLruCache(capacity int) LruCache {
	return &Cache{
		capacity: capacity,
		list:     List{nil},
	}
}
