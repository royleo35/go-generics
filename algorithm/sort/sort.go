// 支持基础类型排序, 移植了sort包堆排序算法
package sort

import (
	"github.com/royleo35/go-generics/types"
	"sort"
)

/*
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}
*/

func less[T types.BuiltIn](data []T, i, j int) bool {
	return data[i] < data[j]
}

func swap[T any](data []T, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func siftDown[T types.BuiltIn](data []T, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(data, first+child, first+child+1) {
			child++
		}
		if !less(data, first+root, first+child) {
			return
		}
		swap(data, first+root, first+child)
		root = child
	}
}

func heapSort[T types.BuiltIn](data []T, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		swap(data, first, first+i)
		siftDown(data, lo, i, first)
	}
}

func Reverse[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		swap(data, i, j)
	}
}

// 移植sort包的堆排序算法
func sortAux[T types.BuiltIn](data []T, desc bool) {
	if len(data) <= 1 {
		return
	}
	heapSort(data, 0, len(data))
	if desc {
		Reverse(data)
	}
}

func Sort[T types.BuiltIn](data []T) {
	sortAux(data, false)
}

func SortDesc[T types.BuiltIn](data []T) {
	sortAux(data, true)
}

func SortWrap[T types.BuiltIn](data []T) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func SortByCopyDef[T types.BuiltIn](data []T) []T {
	s := make([]T, len(data))
	copy(s, data)
	Sort(s)
	return s
}

func SortByCopy[T any](data []T, less func(i, j int) bool) []T {
	s := make([]T, len(data))
	copy(s, data)
	sort.Slice(s, less)
	return s
}
