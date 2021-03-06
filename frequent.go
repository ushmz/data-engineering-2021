package main

import (
	"container/heap"
	"errors"
	"fmt"
)

func MajorityCount(stream []interface{}) (interface{}, error) {
	var majority interface{}
	counter := 0

	for _, v := range stream {
		if majority == nil {
			majority = v
			counter++
		} else if majority == v {
			counter++
		} else {
			counter--
			if counter == 0 {
				majority = nil
			}
		}
		fmt.Println("Item:", majority, "\t Counter: ", counter)
	}

	if majority == nil {
		return nil, errors.New("There is no majority elements")
	}

	return majority, nil
}

func Frequent(stream []interface{}, k int) (interface{}, error) {
	t := map[interface{}]int{}

	for i, v := range stream {
		if (i+1)%k == 0 {
			for key, value := range t {
				if value == 1 {
					delete(t, key)
				}
			}
		}

		if val, ok := t[v]; ok {
			t[v] = val + 1
		} else {
			t[v] = 1
		}

		fmt.Println("Time Stamp:", i, "\t: ", t)
	}

	return t, nil
}

func LossyCounting(stream []interface{}, k int) (interface{}, error) {
	delta := 0
	t := map[interface{}]int{}

	for i, v := range stream {
		if val, ok := t[v]; ok {
			t[v] = val + 1
		} else {
			t[v] = delta + 1
		}

		if (i+1)/k != delta {
			delta = (i + 1) / k
			for key, value := range t {
				if value <= delta {
					delete(t, key)
				}
			}
		}
		fmt.Println("Time Stamp:", i, "\t: ", t)
	}
	return t, nil
}

type Counter struct {
	Item  interface{}
	Count int
}

type StreamHeap []Counter

func (s StreamHeap) Len() int { return len(s) }

func (s StreamHeap) Less(i, j int) bool {
	return s[i].Count < s[j].Count
}

func (s *StreamHeap) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *StreamHeap) Push(item interface{}) {
	*s = append(*s, item.(Counter))
}

func (s *StreamHeap) Pop() interface{} {
	old := *s
	n := len(old)
	item := old[n-1]
	*s = old[0 : n-1]
	return item
}

func SpaceSaving(stream []interface{}, k int) (interface{}, error) {
	t := &StreamHeap{}
	heap.Init(t)

	for i, v := range stream {
		found := false
		for idx, val := range *t {
			if val.Item == v {
				found = true
				(*t)[idx] = Counter{Item: val.Item, Count: val.Count + 1}
			}
		}

		if !found {
			if t.Len() < k {
				heap.Push(t, Counter{Item: v, Count: 1})
			} else {
				min := heap.Pop(t).(Counter)
				heap.Push(t, Counter{Item: v, Count: min.Count + 1})
			}
		}
		fmt.Println("Time Stamp:", i, "\t: ", t)
	}
	return t, nil
}
