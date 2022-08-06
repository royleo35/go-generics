// treemap 使用rb-tree作为底层存储，对外提供hash表

package treemap

import (
	"github.com/royleo35/go-generics/rbtree"
	"github.com/royleo35/go-generics/types"
)

type Pair[K, V any] struct {
	key K
	val V
}

type TreeMap[K, V any] rbtree.RBTree[*Pair[K, V]]

func castFromRBTree[K, V any](t *rbtree.RBTree[*Pair[K, V]]) *TreeMap[K, V] {
	return (*TreeMap[K, V])(t)
}

func (tm *TreeMap[K, V]) castToRBTree() *rbtree.RBTree[*Pair[K, V]] {
	return (*rbtree.RBTree[*Pair[K, V]])(tm)
}

func NewPair[K, V any](key K, val V) *Pair[K, V] {
	return &Pair[K, V]{
		key: key,
		val: val,
	}
}

func (p *Pair[K, V]) Key() K {
	return p.key
}

func (p *Pair[K, V]) Value() V {
	return p.val
}

func defCmpKey[K types.BuiltIn](k1, k2 K) int {
	if k1 > k2 {
		return 1
	} else if k1 < k2 {
		return -1
	}
	return 0
}

//func defCmp[K types.BuiltIn, V any](v1, v2 *Pair[K, V]) int {
//	return defCmpKey(v1.key, v2.key)
//}

func cmp[K, V any](c func(k1, k2 K) int) func(p1, p2 *Pair[K, V]) int {
	return func(p1, p2 *Pair[K, V]) int {
		return c(p1.key, p2.key)
	}
}

func NewDef[K types.BuiltIn, V any]() *TreeMap[K, V] {
	return New[K, V](defCmpKey[K])
}

func NewDefByNodes[K types.BuiltIn, V any](s []*Pair[K, V]) *TreeMap[K, V] {
	tm := New[K, V](defCmpKey[K])
	for _, v := range s {
		tm.Set(v.key, v.val)
	}
	return tm
}

func New[K, V any](compare func(k1, k2 K) int) *TreeMap[K, V] {
	return castFromRBTree(rbtree.NewRBTree[*Pair[K, V]](cmp[K, V](compare)))
}

func NewDefByValues[K types.BuiltIn, V any](values []V, f func(V) K) *TreeMap[K, V] {
	tm := NewDef[K, V]()
	for _, v := range values {
		tm.Set(f(v), v)
	}
	return tm
}

func (tm *TreeMap[K, V]) Set(key K, val V) {
	p := NewPair(key, val)
	tm.AddPair(p)
}

func (tm *TreeMap[K, V]) AddPair(p *Pair[K, V]) {
	tm.castToRBTree().Insert(p)
}

func (tm *TreeMap[K, V]) Delete(key K) {
	var val V
	p := NewPair(key, val)
	tm.deletePair(p)
}

func (tm *TreeMap[K, V]) deletePair(p *Pair[K, V]) {
	tm.castToRBTree().Delete(p)
}

//func (tm *TreeMap[K, V]) Keys() []K {
//	res := slice.Map(tm.castToRBTree().ToSlice(), func(p *Pair[K, V]) K {
//		return p.key
//	})
//	return res
//}
//
//func (tm *TreeMap[K, V]) Values() []V {
//	res := slice.Map(tm.castToRBTree().ToSlice(), func(p *Pair[K, V]) V {
//		return p.val
//	})
//	return res
//}

func (tm *TreeMap[K, V]) ToSlice() []*Pair[K, V] {
	return tm.castToRBTree().ToSlice()
}

func (tm *TreeMap[K, V]) Keys() []K {
	return tm.OrderKeys()
}

func (tm *TreeMap[K, V]) Values() []V {
	return tm.OrderValues()
}

func (tm *TreeMap[K, V]) OrderKeys() []K {
	t := tm.castToRBTree()
	res := make([]K, t.Len())
	for node, i := t.First(), 0; node != nil; node, i = node.Next(), i+1 {
		res[i] = node.Value().key
	}
	return res
}

func (tm *TreeMap[K, V]) OrderValues() []V {
	t := tm.castToRBTree()
	res := make([]V, t.Len())
	for node, i := t.First(), 0; node != nil; node, i = node.Next(), i+1 {
		res[i] = node.Value().val
	}
	return res
}

func (tm *TreeMap[K, V]) Contain(key K) bool {
	_, ok := tm.Get(key)
	return ok
}

func (tm *TreeMap[K, V]) Get(key K) (val V, ok bool) {
	p := NewPair(key, val)
	v := tm.castToRBTree().Find(p)
	if v == nil {
		return val, false
	}
	return v.Value().val, true
}

type Iter[K any, V any] struct {
	node *rbtree.TreeNode[*Pair[K, V]]
}

func newIter[K any, V any](node *rbtree.TreeNode[*Pair[K, V]]) *Iter[K, V] {
	return &Iter[K, V]{node}
}

func (tm *TreeMap[K, V]) First() *Iter[K, V] {
	f := tm.castToRBTree().First()
	if f == nil {
		return nil
	}
	return newIter(f)
}
func (i *Iter[K, V]) Next() *Iter[K, V] {
	n := i.node.Next()
	if n == nil {
		return nil
	}
	return newIter(n)
}

func (i *Iter[K, V]) Pair() *Pair[K, V] {
	return i.node.Value()
}

func (i *Iter[K, V]) KV() (K, V) {
	return i.Key(), i.Value()
}

func (i *Iter[K, V]) Key() K {
	return i.Pair().key
}

func (i *Iter[K, V]) Value() V {
	return i.Pair().val
}
