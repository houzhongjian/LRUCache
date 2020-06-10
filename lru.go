package lru

import (
	"container/list"
	"log"
	"sync"
)

type LRUCache struct {
	num   int
	list  *list.List
	cache map[string]Node
	sync.RWMutex
}

type Node struct {
	Element *list.Element
	Val     interface{}
}

func New(num int) *LRUCache {
	list.New()
	return &LRUCache{
		num:     num,
		list:    list.New(),
		cache:   make(map[string]Node),
		RWMutex: sync.RWMutex{},
	}
}

func (lc *LRUCache) Set(key string, val interface{}) {
	defer lc.Unlock()
	lc.Lock()

	//判断当前key是否已经存在.
	node, ok := lc.cache[key]
	if ok {
		lc.cache[key] = Node{
			Val:     val,
			Element: node.Element,
		}
		return
	}

	log.Println("添加元素:", key)
	if lc.list.Len() >= lc.num {
		//获取列表最后的一个元素
		backElement := lc.list.Back()
		expiredKey := backElement.Value.(string)
		log.Println("删除列表最后一个元素:", expiredKey)
		delete(lc.cache, expiredKey)
		lc.list.Remove(backElement)
	}

	element := lc.list.PushFront(key)
	lc.cache[key] = Node{
		Val:     val,
		Element: element,
	}
}

func (lc *LRUCache) Get(key string) interface{} {
	defer lc.RUnlock()
	lc.RLock()

	node, ok := lc.cache[key]
	if !ok {
		return nil
	}

	//将当前元素添加到list头部.
	lc.list.MoveToFront(node.Element)
	return node.Val
}

func (lc *LRUCache) Delete(key string) {
	defer lc.Unlock()
	lc.Lock()

	node, ok := lc.cache[key]
	if !ok {
		return
	}

	delete(lc.cache, key)
	lc.list.Remove(node.Element)
}
