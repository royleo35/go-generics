package hashmap

import (
	"fmt"
	"github.com/royleo35/go-generics/glist"
	"github.com/royleo35/go-generics/tools"
	"unsafe"
)

type pair[K, V any] struct {
	key K
	val V
}

type HashMap[K, V any] struct {
	size      int
	bucketIdx int
	fmtKey    func(K) string // 用户提供将key转为string的函数,同时用于hash和比较
	oldBucket []*glist.List[*pair[K, V]]
	newBucket []*glist.List[*pair[K, V]]
	// rehash进度, 下一个需要rehash的bucket索引， -1 代表没在rehash
	rehashingIdx int
	// grow 标识rehash过程是扩容还是缩容
	grow bool
}

var (
	bucketSize = [...]int{
		2, 5, 11, 23, 47, 97, 197, 397, 797, 1597, 3203, 6421, 12853, 25717, 51437, 102877, 205759, 411527, 823117, 1646237,
		3292489, 6584983, 13169977, 26339969, 52679969, 105359939, 210719881, 421439783, 842879579}
	//bucketSize = [...]int{
	//	2, 4, 8, 16, 32, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20,
	//	1 << 21, 1 << 22, 1 << 23, 1 << 24, 1 << 25, 1 << 26, 1 << 27, 1 << 28, 1 << 29}
	bucketSizeLen = len(bucketSize)
)

const (
	maxFactor      = 1.0
	minFactor      = 0.1
	onceRehashSlot = 3
)

func defFmtKey[K comparable](k K) string {
	return fmt.Sprintf("%v", k)
}

func findBuckSizeIdx(size int) int {
	for i := 0; i < bucketSizeLen; i++ {
		if size <= bucketSize[i] {
			return i
		}
	}
	return len(bucketSize) - 1
}

func assert(cond bool, errMsg string) {
	if !cond {
		panic(errMsg)
	}
}

func lastIdx(idx int) int {
	assert(idx >= 0 && idx <= bucketSizeLen-1, "index error")
	if idx == 0 {
		return 0
	}
	return idx - 1
}

func nextIdx(idx int) int {
	assert(idx >= 0 && idx <= bucketSizeLen-1, "index error")
	if idx == bucketSizeLen-1 {
		return idx
	}
	return idx + 1
}

func NewHashMap[K, V any](size int, fmtKey func(K) string) *HashMap[K, V] {
	if size < 0 {
		panic("size must >= 0")
	}
	bucketIdx := findBuckSizeIdx(size)
	m := &HashMap[K, V]{
		size:         0,
		bucketIdx:    bucketIdx,
		fmtKey:       fmtKey,
		oldBucket:    allocAux[*pair[K, V]](bucketSize[bucketIdx]),
		rehashingIdx: -1,
	}
	return m
}

func NewDefHashMap[K comparable, V any](size int) *HashMap[K, V] {
	return NewHashMap[K, V](size, defFmtKey[K])
}

func NewWithValues[K comparable, V, T any](s []T, f func(e T) (K, V)) *HashMap[K, V] {
	m := NewDefHashMap[K, V](len(s))
	for _, v := range s {
		key, val := f(v)
		m.Set(key, val)
	}
	return m
}

func allocAux[T any](size int) []*glist.List[T] {
	res := make([]*glist.List[T], size)
	for i := range res {
		res[i] = glist.New[T]()
	}
	return res
}

func rotateShiftRight(val uint64, bits int) uint64 {
	return (val >> bits) | (val << (64 - bits))
}

func hash(key string) uint64 {
	res := uint64(0)
	l := len(key)
	p := unsafe.Pointer(&key)
	l8 := l >> 3
	data := *(*[]byte)(unsafe.Pointer(uintptr(p) + uintptr(l8<<3)))
	for i := 0; i < l8; i++ {
		res |= rotateShiftRight(*(*uint64)(unsafe.Pointer(uintptr(p) + uintptr(i<<3))), 8)
	}
	rest := l & 7
	for i := 0; i < rest; i++ {
		res |= rotateShiftRight(uint64(data[i]), 4)
	}
	return res
}

func (hm *HashMap[K, V]) hashIdx(key K, bucket int) int {
	if bucket == 0 {
		return 0
	}
	idx := hash(hm.fmtKey(key)) % uint64(bucket)
	return int(idx)
}

func (hm *HashMap[K, V]) BucketCount() int {
	return bucketSize[hm.bucketIdx]
}

func (hm *HashMap[K, V]) NextBucketCount() int {
	return bucketSize[nextIdx(hm.bucketIdx)]
}

func (hm *HashMap[K, V]) LastBucketCount() int {
	return bucketSize[lastIdx(hm.bucketIdx)]
}

func (hm *HashMap[K, V]) nextRehashIdx() int {
	if hm.grow {
		return nextIdx(hm.bucketIdx)
	}
	return lastIdx(hm.bucketIdx)
}

func (hm *HashMap[K, V]) rehashBucketCount() int {
	if hm.grow {
		return hm.NextBucketCount()
	}
	return hm.LastBucketCount()
}

func (hm *HashMap[K, V]) rehashing() bool {
	return hm.rehashingIdx != -1
}

func (hm *HashMap[K, V]) factor() float64 {
	return float64(hm.size) / float64(hm.BucketCount())
}

// 做一些rehash的工作，一次rehash3个bucket
func (hm *HashMap[K, V]) rehash() {
	if !hm.rehashing() {
		if hm.factor() < minFactor && hm.LastBucketCount() != hm.BucketCount() { // 缩容
			hm.newBucket = allocAux[*pair[K, V]](hm.LastBucketCount())
			hm.rehashingIdx = 0
			hm.grow = false
		} else if hm.factor() > maxFactor && hm.NextBucketCount() != hm.BucketCount() { // 扩容
			hm.newBucket = allocAux[*pair[K, V]](hm.NextBucketCount())
			hm.rehashingIdx = 0
			hm.grow = true
		} else {
			return
		}
	}
	bucket := hm.BucketCount()
	for i := 0; hm.rehashingIdx < bucket && i < onceRehashSlot; hm.rehashingIdx, i = hm.rehashingIdx+1, i+1 {
		slot := hm.oldBucket[hm.rehashingIdx]
		if slot.Len() == 0 {
			continue
		}
		for e := slot.Front(); e != nil; e = e.Next() {
			hm.set(e.Value.key, e.Value.val, false)
		}
		hm.oldBucket[hm.rehashingIdx] = nil
	}
	// finished rehash, swap bucket
	if hm.rehashingIdx == bucket {
		hm.oldBucket = hm.newBucket
		hm.newBucket = nil
		hm.bucketIdx = hm.nextRehashIdx()
		hm.rehashingIdx = -1
	}
}

func (hm *HashMap[K, V]) find(key K, old bool) (val V, ok bool) {
	e := hm.findElem(key, old)
	if e == nil {
		return
	}
	return e.Value.val, true
}

func (hm *HashMap[K, V]) findElem(key K, old bool) *glist.Element[*pair[K, V]] {
	var (
		b   []*glist.List[*pair[K, V]]
		idx int
	)
	if old {
		b = hm.oldBucket
		idx = hm.hashIdx(key, hm.BucketCount())
	} else {
		b = hm.newBucket
		idx = hm.hashIdx(key, hm.rehashBucketCount())
	}
	if len(b) == 0 || b[idx] == nil {
		return nil
	}
	for e := b[idx].Front(); e != nil; e = e.Next() {
		if hm.fmtKey(key) == hm.fmtKey(e.Value.key) {
			return e
		}
	}
	return nil
}

func (hm *HashMap[K, V]) getElem(key K) *glist.Element[*pair[K, V]] {
	// 先从旧bucket找
	e := hm.findElem(key, true)
	if e != nil {
		return e
	}
	// 如果正在rehash，从新bucket找
	if hm.rehashing() {
		return hm.findElem(key, false)
	}
	return nil
}

func (hm *HashMap[K, V]) Get(key K) (val V, ok bool) {
	e := hm.getElem(key)
	if e == nil {
		return
	}
	return e.Value.val, true
}

func (hm *HashMap[K, V]) set(key K, val V, old bool) {
	if old {
		idx := hm.hashIdx(key, hm.BucketCount())
		slot := hm.oldBucket[idx]
		slot.PushFront(&pair[K, V]{key, val})
	} else {
		idx := hm.hashIdx(key, hm.rehashBucketCount())
		slot := hm.newBucket[idx]
		slot.PushFront(&pair[K, V]{key, val})
	}
}

func (hm *HashMap[K, V]) Set(key K, val V) {
	e := hm.getElem(key)
	if e != nil {
		e.Value.key = key
		e.Value.val = val
		return
	}
	if !hm.rehashing() {
		hm.set(key, val, true)
	} else {
		hm.set(key, val, false)
	}
	hm.size++
	hm.rehash()
}

func (hm *HashMap[K, V]) Size() int {
	return hm.size
}

type Iter[K, V any] struct {
	hm *HashMap[K, V]
	// 用于处理边界情况，第一次迭代和旧bucket到新bucket切换时从currSlot开始扫描，其他情况从currSlot+1开始扫描
	iterSlot      bool
	currSlot      int
	currElem      *glist.Element[*pair[K, V]]
	currBucketOld bool
}

func (it *Iter[K, V]) KV() (key K, val V) {
	return it.Key(), it.Value()
}
func (it *Iter[K, V]) Key() K {
	return it.currElem.Value.key
}

func (it *Iter[K, V]) Value() V {
	return it.currElem.Value.val
}

func (it *Iter[K, V]) Next() *Iter[K, V] {
	if it.hm.size == 0 {
		return nil
	}
	var (
		b        []*glist.List[*pair[K, V]]
		l        int
		nextElem *glist.Element[*pair[K, V]]
	)
	// first
	if it.currSlot == 0 && it.currElem == nil {
		it.iterSlot = false
		goto search
	}
	nextElem = it.currElem.Next()
	if nextElem != nil {
		it.currElem = nextElem
		return it
	}
	// not rehashing, no more elem
	if !it.hm.rehashing() && it.currSlot == it.hm.BucketCount()-1 {
		return nil
	}
	// rehashing, no more elem
	if it.hm.rehashing() && it.currSlot == it.hm.rehashBucketCount()-1 {
		return nil
	}
	//// switch old to new
	//if it.currSlot == it.hm.BucketCount() && it.currBucketOld == true && it.hm.rehashing() {
	//	it.currSlot = 0
	//	it.currBucketOld = false
	//	it.iterSlot = false
	//}
search:
	if it.currBucketOld {
		b = it.hm.oldBucket
		l = it.hm.BucketCount()
	} else {
		b = it.hm.newBucket
		l = it.hm.rehashBucketCount()
	}
	// 首次迭代和bucket切换时从当前slot开始扫描
	// 其他情况，上面nextEle==nil,说明当前slot没有元素了，从下一个slot开始扫描
	for i := tools.IfThen(it.iterSlot, it.currSlot+1, it.currSlot); i < l; i++ {
		if b[i] == nil || b[i].Len() == 0 {
			continue
		}
		it.currElem = b[i].Front()
		it.currSlot = i
		it.iterSlot = true
		return it
	}
	// switch old to new
	if it.hm.rehashing() && it.currBucketOld {
		it.currSlot = 0
		it.currBucketOld = false
		it.iterSlot = false
		goto search
	}
	return nil
}

func (hm *HashMap[K, V]) First() *Iter[K, V] {
	if hm.size == 0 {
		return nil
	}
	it := &Iter[K, V]{
		hm:            hm,
		currSlot:      0,
		currElem:      nil,
		currBucketOld: true,
	}
	return it.Next()
}

func (hm *HashMap[K, V]) Delete(key K) {
	e := hm.getElem(key)
	if e != nil {
		e.List().Remove(e) // 杀死自己 哈哈
		hm.size--
	}
	hm.rehash()
}

func (hm *HashMap[K, V]) Keys() []K {
	res := make([]K, hm.size)
	for i, it := 0, hm.First(); it != nil; i, it = i+1, it.Next() {
		res[i] = it.Key()
	}
	return res
}

func (hm *HashMap[K, V]) Values() []V {
	res := make([]V, hm.size)
	for i, it := 0, hm.First(); it != nil; i, it = i+1, it.Next() {
		res[i] = it.Value()
	}
	return res
}
