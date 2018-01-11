package lru

import (
	".."
)

type Lru struct {
	cacher.Cacher
}

func New() *Lru {
	lru :=new(Lru)
	return lru;
}