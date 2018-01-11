package cacher

type Cacher interface {
	Add(key Key,val interface{})
	Get(key Key) (val interface{},ok bool)
	Remove(key Key)
	Clear()
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key interface{}

