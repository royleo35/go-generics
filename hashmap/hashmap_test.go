package hashmap

import (
	"fmt"
	"github.com/royleo35/go-generics/algorithm/sort"
	"github.com/royleo35/go-generics/gmap"
	"github.com/royleo35/go-generics/hputil"
	"github.com/royleo35/go-generics/tools"
	"math/rand"
	"testing"
)

func TestNewDefHashMap(t *testing.T) {
	hm := NewDefHashMap[int64, int64](0)
	// set
	for i := int64(0); i < 1e4; i++ {
		hm.Set(i, i+1)
		if hm.rehashing() {
			print(i)
			break
		}
	}
	// average
	maxSlot := 0
	for i := 0; i < hm.BucketCount(); i++ {
		if hm.oldBucket[i] == nil {
			continue
		}
		if hm.oldBucket[i].Len() > maxSlot {
			maxSlot = hm.oldBucket[i].Len()
		}

		if hm.newBucket[i] == nil {
			continue
		}
		if hm.newBucket[i].Len() > maxSlot {
			maxSlot = hm.newBucket[i].Len()
		}
	}
	print(maxSlot)
	// get
	for i := int64(0); i < int64(hm.size); i++ {
		val, ok := hm.Get(i)
		tools.Assert(ok && val == i+1)
	}

	// iter
	for it := hm.First(); it != nil; it = it.Next() {
		fmt.Println("key:", it.currElem.Value.key, "val:", it.currElem.Value.val)
	}

	// delete
	for i := int64(0); i < int64(hm.size); i++ {
		hm.Delete(i)
		_, ok := hm.Get(i)
		tools.Assert(!ok)
	}
}

func TestFunction(t *testing.T) {
	hm := NewDefHashMap[int, int](0)
	m := map[int]int{}
	const mapSize = 1e2
	// set get
	for i := 0; i < mapSize; i++ {
		key := rand.Int()
		hm.Set(key, i+1)
		m[key] = i + 1
		tools.Assert(hm.Size() == len(m))
		v1, ok1 := hm.Get(i)
		v2, ok2 := m[i]
		tools.Assert(v1 == v2 && ok1 == ok2)
	}
	// iter keys values
	k1 := hm.Keys()
	k2 := gmap.Keys(m)
	tools.Assert(hputil.DeepEqual(sort.SortByCopyDef(k1), sort.SortByCopyDef(k2)))
	v1 := hm.Values()
	v2 := gmap.Values(m)
	tools.Assert(hputil.DeepEqual(sort.SortByCopyDef(v1), sort.SortByCopyDef(v2)))

	// contain each other
	for it := hm.First(); it != nil; it = it.Next() {
		key := it.Key()
		val := it.Value()
		v2, ok := m[key]
		tools.Assert(ok && val == v2)
	}
	for k, v := range m {
		v2, ok := hm.Get(k)
		tools.Assert(ok && v == v2)
	}

	// delete get
	for i := 0; i < mapSize; i++ {
		hm.Delete(i)
		delete(m, i)
		tools.Assert(hm.Size() == len(m))
		v1, ok1 := hm.Get(i)
		v2, ok2 := m[i]
		tools.Assert(v1 == v2 && ok1 == ok2)
	}

}

var randKeys = func() []int {
	res := make([]int, 1e5)
	for i := 0; i < len(res); i++ {
		res[i] = rand.Int()
	}
	return res
}()

// benckmark结果，key为int时性能只有map的1/10
func BenchmarkMapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i, v := range randKeys {
			m[v] = i
		}
	}
}

func BenchmarkHashMapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hm := NewDefHashMap[int, int](0)
		for i, v := range randKeys {
			hm.Set(v, i)
		}
	}
}

func BenchmarkMapSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i, v := range randKeys {
			m[v] = i
			val, ok := m[v]
			_, _ = val, ok
		}
	}
}

func BenchmarkHashMapSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hm := NewDefHashMap[int, int](0)
		for i, v := range randKeys {
			hm.Set(v, i)
			val, ok := hm.Get(v)
			_, _ = val, ok
		}
	}
}
