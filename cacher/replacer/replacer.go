package replacer

type replacer interface {
	replace() []interface{}
	remove(keys []interface{})
	add(key interface{},arguments ...interface{})
}