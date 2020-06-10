package lru

import (
	"container/list"
	"sync"
)

type cache struct {
	sync.RWMutex
	data map[string]Node
}

func (c *cache) add(key string, val interface{}, ele *list.Element) {
	defer c.Unlock()
	c.Lock()
	c.data[key] = Node{
		Element: ele,
		Val:     val,
	}
}

func (c *cache) get(key string) *Node {
	defer c.RUnlock()
	c.RLock()

	node, ok := c.data[key]
	if !ok {
		return nil
	}

	return &node
}

func (c *cache) remove(key string) *list.Element {
	defer c.Unlock()
	c.Lock()

	node, ok := c.data[key]
	if !ok {
		return nil
	}

	delete(c.data, key)

	return node.Element
}