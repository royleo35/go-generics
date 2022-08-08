package myslice

import (
	"github.com/royleo35/go-generics/tools"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestNextPower(t *testing.T) {
	s := nextPower(10)
	s1 := nextPower(127)
	s2 := nextPower(1025)
	s3 := nextPower(0xff003244)
	s4 := nextPower(0x1ffffffff)
	assert(s == 16, "ops")
	assert(s1 == 128, "ops")
	assert(s2 == 2048, "ops")
	assert(s3 == (1<<32), "ops")
	assert(s4 == (1<<33), "ops")

}

func TestAppend(t *testing.T) {
	s := New[int](0, 0)
	s1 := s.Append(10, 20, 30)
	assert(s1.data == s.data && s1.Len() == 3 && s1.At(0) == 10 && s1.At(1) == 20 && s.At(2) == 30, "ops")
	func() {
		defer func() {
			if e := recover(); e != nil {
				assert(strings.Contains(e.(string), "index"), "ouch")
			}
		}()
		_ = s1.At(3)
	}()
	func() {
		defer func() {
			if e := recover(); e != nil {
				assert(strings.Contains(e.(string), "index"), "ouch")
			}
		}()
		_ = s1.At(-1)
	}()

}

func TestFor(t *testing.T) {
	b1 := []byte{1, 2, 3, 4, 5, 6, 7}
	s1 := NewWithValues(b1...)
	for i := 0; i < s1.Len(); i++ {
		assert(s1.At(i) == b1[i], "")
	}
	i2 := []int64{1, 2, 5, 6, 9}
	s2 := NewWithValues(i2...)
	for i := 0; i < s2.Len(); i++ {
		assert(s2.At(i) == i2[i], "")
	}
	b22 := bytes(i2)
	s22 := s2.bytes()
	assert(reflect.DeepEqual(b22, s22), "ops")
}

func TestSlice(t *testing.T) {
	i2 := []int64{1, 2, 5, 6, 9}
	s2 := NewWithValues(i2...)
	s := s2.Slice(0, 2).ToSlice()
	assert(reflect.DeepEqual(s, []int64{1, 2}), "")
	assert(reflect.DeepEqual(s2.Slice(0, 5).ToSlice(), i2), "")
	func() {
		defer func() {
			if e := recover(); e != nil {
				assert(strings.Contains(e.(string), "be in"), "ouch")
			}
		}()
		_ = s2.Slice(-1, 2)
	}()
	func() {
		defer func() {
			if e := recover(); e != nil {
				assert(strings.Contains(e.(string), "be in"), "ouch")
			}
		}()
		_ = s2.Slice(2, 10)
	}()
}

func mockData(n int) ([]byte, []int, []*string) {
	bs := make([]byte, n)
	ints := make([]int, n)
	strs := make([]*string, n)
	for i := 0; i < n; i++ {
		d := rand.Int()
		bs[i] = byte(d & 0xff)
		ints[i] = d
		strs[i] = tools.PtrOf(strconv.FormatInt(int64(i), 10))
	}
	return bs, ints, strs
}

var (
	bs1, ints1, strs1 = mockData(1000)
	mybs1             = NewWithValues(bs1)
	myints1           = NewWithValues(ints1)
	mystrs1           = NewWithValues(strs1)
	bs2, ints2, strs2 = mockData(1e5)
	mybs2             = NewWithValues(bs1)
	myints2           = NewWithValues(ints1)
	mystrs2           = NewWithValues(strs1)
)

func BenchmarkMySliceSet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s1 := New[byte](len(bs1), len(bs1))
		for j, v := range bs1 {
			s1.Set(j, v)
		}
		s2 := New[byte](len(bs2), len(bs2))
		for j, v := range bs2 {
			s2.Set(j, v)
		}
		s3 := New[int](len(ints1), len(ints1))
		for j, v := range ints1 {
			s3.Set(j, v)
		}
		s4 := New[int](len(ints2), len(ints2))
		for j, v := range ints2 {
			s4.Set(j, v)
		}
		s5 := New[*string](len(strs1), len(strs1))
		for j, v := range strs1 {
			s5.Set(j, v)
		}
		s6 := New[*string](len(strs2), len(strs2))
		for j, v := range strs2 {
			s6.Set(j, v)
		}
	}
}

func BenchmarkSliceSet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s1 := make([]byte, len(bs1))
		for j, v := range bs1 {
			s1[j] = v
		}
		s2 := make([]byte, len(bs2))
		for j, v := range bs2 {
			s2[j] = v
		}
		s3 := make([]int, len(ints1))
		for j, v := range ints1 {
			s3[j] = v
		}
		s4 := make([]int, len(ints2))
		for j, v := range ints2 {
			s4[j] = v
		}
		s5 := make([]*string, len(strs1))
		for j, v := range strs1 {
			s5[j] = v
		}
		s6 := make([]*string, len(strs2))
		for j, v := range strs2 {
			s6[j] = v
		}
	}
}

func BenchmarkMySliceGet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < mybs1.Len(); i++ {
			_ = mybs1.At(i)
			_ = mybs2.At(i)
			_ = myints1.At(i)
			_ = myints2.At(i)
			_ = mystrs1.At(i)
			_ = mystrs2.At(i)
		}
	}
}

func BenchmarkSliceGet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(bs1); i++ {
			_ = bs1[i]
			_ = bs2[i]
			_ = ints1[i]
			_ = ints2[i]
			_ = strs1[i]
			_ = strs2[i]
		}
	}
}

func BenchmarkMySliceAppend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New[int](0, 0)
		s2 := New[int](0, 0)
		for _, v := range ints1 {
			s = s.Append(v)
		}
		for _, v := range ints2 {
			s2 = s2.Append(v)
		}
	}
}

func BenchmarkSliceAppend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		s2 := make([]int, 0)
		for _, v := range ints1 {
			s = append(s, v)
		}
		for _, v := range ints2 {
			s2 = append(s2, v)
		}
	}
}

func BenchmarkMySliceAppend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New[int](0, 0)
		for _, v := range ints1 {
			s = s.Append(v)
		}
		s = s.Append(ints2...)
	}
}

func BenchmarkSliceAppend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for _, v := range ints1 {
			s = append(s, v)
		}
		s = append(s, ints2...)
	}
}
