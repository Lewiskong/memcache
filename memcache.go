package memcache

import (
	"./cacher"
	"./cacher/sortTree"
	"./cacher/lru"
	"sync"
	"time"
)

type CacherType int32
const (
	ExpireTimeCacher CacherType=1
	LruCacher CacherType=2

	DefaultCacher = ExpireTimeCacher
)

type memCache struct {
	name string

	mu sync.RWMutex
	nhit,nget int64

	cacher cacher.Cacher
	ctype CacherType
}



func NewMemCache(ctype CacherType) *memCache{
	cache := new(memCache)
	switch ctype {
	case ExpireTimeCacher:
		cache.cacher=sortTree.New()
	case LruCacher:
		cache.cacher=lru.New()
	default:
		panic("unknown cacher type")
	}

	return cache
}



func (cache *memCache) Set(key cacher.Key,val interface{},expireRemove ...time.Duration ){
	switch cache.ctype {
	case ExpireTimeCacher:
		var expireTime,removeTime time.Duration
		if len(expireRemove)==1 {
			expireTime,removeTime = expireRemove[0],expireRemove[0]+time.Minute
		}else if len(expireRemove)==2 {
			expireTime,removeTime = expireRemove[0],expireRemove[0] + expireRemove[1]
		}else {
			panic("invalid parameter when call set func")
		}
		cache.setSortTree(key,val,expireTime,removeTime)
	case LruCacher:
		cache.setLru(key,val)
	default:
		panic("unknown cacher type")
	}
}

func (cache *memCache) setSortTree(key cacher.Key,val interface{},expireTime,removeTime time.Duration){

}

func (cache *memCache) setLru(key cacher.Key,val interface{}){

}

func (cache *memCache) Get(key cacher.Key) (interface{},bool){
	return nil,true
}





