package replacer

import (
	"sync"
	"sync/atomic"
)

const (
	nodeBufSize = 1024
)

// 只有左枝的二叉排序树
// 场景特点 无需用AVL树
type tree struct {
	mutex sync.Mutex
	root  *node

	nodeBuf chan *node
	nodeCnt int64 // 缓存节点的数量

	replacer
}

type node struct {
	key   interface{}
	value int // 过期时间
	left  *node
	right *node
}

func (t *tree) replace() []interface{} {
	return nil
}

func (t *tree) remove(keys []interface{}) {

}

func (t *tree) add(key interface{}, arguments ...interface{}) {
	// 参数检查
	if len(arguments) != 1 {
		return
	}
	expire, ok := arguments[0].(int)
	if !ok {
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 缓存键相等，更新过期时间
	if t.root != nil && t.root.key == key {
		t.root.value = expire
		return
	}
	// 比根节点大，重新创建根节点
	if t.root == nil || expire > t.root.value {
		addNode := t.newNode(key, expire)
		addNode.left = t.root
		t.root = addNode
		return
	}

	curNode := t.root
	for {
		if key == curNode.key {
			curNode.value = expire
			return
		}
		if expire < curNode.value {
			if curNode.left == nil {
				addNode := t.newNode(key, expire)
				curNode.left = addNode
				return
			}
			curNode = curNode.left
			continue
		}

		if expire > curNode.value {
			if curNode.right == nil {
				addNode := t.newNode(key, expire)
				curNode.right = addNode
				return
			}
			curNode = curNode.right
			continue
		}
	}

}

func (t *tree) newNode(key interface{}, val int) (n *node) {
	select {
	case n = <-t.nodeBuf:
	default:
		n = new(node)
	}
	n.key = key
	n.value = val
	n.left = nil
	n.right = nil
	atomic.AddInt64(&t.nodeCnt, 1)
	return
}

func newTree() *tree {
	tree := new(tree)
	tree.nodeBuf = make(chan *node, nodeBufSize)
	return tree
}
