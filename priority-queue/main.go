package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	// Добавляем время окончания первой встречи
	heap.Push(h, intervals[0][1])

	// Проходим по всем встречам
	for i := 1; i < len(intervals); i++ {
		// Если начало следующей встречи >= чем минимальное время окончания
		if intervals[i][0] >= (*h)[0] {
			heap.Pop(h) // Переиспользуем комнату
		}
		// Добавляем новое время окончания встречи
		heap.Push(h, intervals[i][1])
	}

	// Длина кучи — это минимальное количество необходимых комнат
	fmt.Println(len(*h))
}
