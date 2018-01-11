package sortTree

import (
	".."
	"sync"
)

// 利用二叉排序树根据过期时间来进行淘汰
//

type SortTree struct {
	mutex   sync.Mutex
	root    *Node
	keyCnt  int
	valCnt  int
	nodeBuf chan *Node

	cacher.Cacher
}

type Node struct {
	key   cacher.Key
	vals  []interface{}
	left  *Node
	right *Node
}

func New() *SortTree {
	tree :=new(SortTree)
	tree.nodeBuf = make(chan *Node)
	return tree
}

func (tree *SortTree) Add(key cacher.Key, val interface{}) {

}

func (tree *SortTree) Get(key cacher.Key) (val interface{}, ok bool) {
	return nil, true
}

func (tree *SortTree) Remove(key cacher.Key) {

}

func (tree *SortTree) Clear() {

}
