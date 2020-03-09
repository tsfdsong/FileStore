package main

import (
	"fmt"
	"sync"
	"time"
)

var instance *LRUCache
var once sync.Once

type lruObj struct {
	time int64
	data string
}

//LRUCache ...
type LRUCache struct {
	timeOut int64
	list    map[string]*lruObj
	l       sync.Mutex
}

//newLRUCache ...
func newLRUCache(num int, timeOut int64) (res *LRUCache) {
	res = &LRUCache{
		timeOut: timeOut,
		list:    make(map[string]*lruObj, num),
	}

	go func() {
		timer := time.NewTicker(time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				res.l.Lock()
				for k, v := range res.list {
					if time.Now().Unix() > v.time {
						Debugf("contract address cache timeout,user=%v store contract address=%v", k, v)
						delete(res.list, k)
					}
				}
				res.l.Unlock()
			}

		}
	}()

	return res
}

//Put ...
func (lru *LRUCache) Put(key, value string) {
	lru.l.Lock()
	defer lru.l.Unlock()

	if _, ok := lru.list[key]; !ok {
		obj := &lruObj{
			time: time.Now().Unix() + lru.timeOut,
			data: value,
		}

		lru.list[key] = obj

		Debugf("contract address cache put success, user=%v store contract address=%v", key, value)
	}
}

//Get ...
func (lru *LRUCache) Get(key string) (string, error) {
	lru.l.Lock()
	defer lru.l.Unlock()

	obj, ok := lru.list[key]
	if !ok {
		return "", fmt.Errorf("Get obj %v not exist", key)
	}

	if obj.data == "" {
		return "", fmt.Errorf("Get data %v not exist", key)
	}

	Debugf("contract address cache get success, user=%v address=%v", key, obj.data)

	return obj.data, nil
}

//GetInstance ...
func GetInstance() *LRUCache {
	once.Do(func() {
		instance = newLRUCache(100, 24*60)
	})
	return instance
}
