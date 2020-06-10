package main

import (
	"log"
	"pandaschool.net/demo/lru"
)

func main() {
	cache := lru.New(2)
	cache.Set("name", "zhangsan")
	cache.Set("age", 10)
	log.Println(cache.Get("name"))
	cache.Set("sex", "男")
	log.Println(cache.Get("name"))
	log.Println(cache.Get("sex"))
	log.Println(cache.Get("name"))
	cache.Set("age", 10)
	cache.Set("sex", "男")
	cache.Set("sex1", "男")
}
