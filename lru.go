package lru

import (
	"container/list"
	"log"
)

type LRUCache struct {
	num   int
	list  *list.List
	cache *cache
}

type Node struct {
	Element *list.Element
	Val     interface{}
}

func New(num int) *LRUCache {
	list.New()
	return &LRUCache{
		num:  num,
		list: list.New(),
		cache: &cache{
			data: make(map[string]Node),
		},
	}
}

func (lc *LRUCache) Set(key string, val interface{}) {
	log.Println("添加元素:", key)
	if lc.list.Len() >= lc.num {
		//获取列表最后的一个元素
		backElement := lc.list.Back()
		expiredKey := backElement.Value.(string)
		log.Println("删除列表最后一个元素:", expiredKey)
		lc.Delete(expiredKey)
	}

	element := lc.list.PushFront(key)
	lc.cache.add(key, val, element)
}

func (lc *LRUCache) Get(key string) interface{} {
	node := lc.cache.get(key)
	if node != nil {
		//将当前元素添加到list头部.
		lc.list.MoveToFront(node.Element)
		return node.Val
	}

	return nil
}

func (lc *LRUCache) Delete(key string) {
	ele := lc.cache.remove(key)
	if ele != nil {
		lc.list.Remove(ele)
	}
}
