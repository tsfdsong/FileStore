package main

import (
	"fmt"
	"sync"
	"time"
)

var objInstance *ObjectCache
var objOnce sync.Once

type userObj struct {
	time        int64
	Objs        map[string]*Objects
	ContainerID string
}

//ObjectCache ...
type ObjectCache struct {
	timeOut int64
	list    map[string]*userObj
	l       sync.Mutex
}

//newLRUCache ...
func newObjectCache(num int, timeOut int64) (res *ObjectCache) {
	res = &ObjectCache{
		timeOut: timeOut,
		list:    make(map[string]*userObj, num),
	}

	go func() {
		timer := time.NewTicker(10 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				res.l.Lock()
				for k, v := range res.list {
					if time.Now().Unix() > v.time {
						Debugf("user=%v object cache timeout,indexhash=%v", k, v.Objs)
						err := CloseChannel(v.ContainerID)
						if err != nil {
							Errorf("Close channel failed, user=%v container=%v", k, v.ContainerID)
						}
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
func (obj *ObjectCache) Put(user, indexhash string, in *Objects) {
	obj.l.Lock()
	defer obj.l.Unlock()

	item, ok := obj.list[user]
	if !ok {
		temp := new(userObj)
		temp.time = time.Now().Unix() + obj.timeOut
		temp.Objs = make(map[string]*Objects)
		temp.Objs[indexhash] = in

		obj.list[user] = temp
	} else {
		item.Objs = make(map[string]*Objects)
		item.Objs[indexhash] = in
	}

	Debugf("object cache put success, user=%v indexhash=%v indexfile=%v", user, indexhash, in.String())
}

//PutContainer ...
func (obj *ObjectCache) PutContainer(user, id string) {
	obj.l.Lock()
	defer obj.l.Unlock()

	item, ok := obj.list[user]
	if !ok {
		temp := new(userObj)
		temp.time = time.Now().Unix() + obj.timeOut
		temp.Objs = make(map[string]*Objects)
		temp.ContainerID = id
		obj.list[user] = temp
	} else {
		item.ContainerID = id
	}

	Debugf("object cache put success, user=%v container=%v", user, id)
}

//Get ...
func (obj *ObjectCache) Get(user, indexhash string) (*Objects, error) {
	obj.l.Lock()
	defer obj.l.Unlock()

	userObj, ok := obj.list[user]
	if !ok {
		return nil, fmt.Errorf("user=%v get indexhash=%v not exist", user, indexhash)
	}

	if res, ok := userObj.Objs[indexhash]; ok {
		Debugf("object cache get success, user=%v indexhash=%v indexfile=%v", user, indexhash, res.String())
		return res, nil
	}

	return nil, fmt.Errorf("user=%v get indexhash=%v object not exist", user, indexhash)
}

//GetContainer ...
func (obj *ObjectCache) GetContainer(user string) (string, error) {
	obj.l.Lock()
	defer obj.l.Unlock()

	userObj, ok := obj.list[user]
	if !ok {
		return "", fmt.Errorf("get container useuser=%v not exist", user)
	}

	if userObj.ContainerID == "" {
		return "", fmt.Errorf("get container useuser=%v empty", user)
	}

	return userObj.ContainerID, nil
}

//GetObjectInstance ...
func GetObjectInstance() *ObjectCache {
	objOnce.Do(func() {
		objInstance = newObjectCache(100, 24*60*60)
	})
	return objInstance
}
