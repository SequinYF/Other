package main

import "container/list"

type LRUCache struct {
	capacity int
	head *list.List
	node map[int]*list.Element
}

type element struct {
	key, value int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head: list.New(),
		node: make(map[int]*list.Element)}
}


func (this *LRUCache) Get(key int) int {
	if elem, ok := this.node[key]; ok {
		this.head.MoveToFront(elem)
		return elem.Value.(*element).value
	}
	return 0
}


func (this *LRUCache) Put(key int, value int)  {

	if elem, ok := this.node[key]; ok {
		this.head.MoveToFront(elem)
		elem.Value.(*element).value = value
	} else {
		newelem := this.head.PushFront(&element{key, value})
		this.node[key] = newelem
	}
	
	if this.head.Len() > this.capacity {
		lastelem := this.head.Back()
		cachenode := lastelem.Value.(*element)
		delete(this.node, cachenode.key)
		this.head.Remove(lastelem)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */