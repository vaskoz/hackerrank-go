package main

import (
	"container/heap"
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

type MinHeap []int

func (min MinHeap) Len() int           { return len(min) }
func (min MinHeap) Less(i, j int) bool { return min[i] < min[j] }
func (min MinHeap) Swap(i, j int)      { min[i], min[j] = min[j], min[i] }

func (min *MinHeap) Push(x interface{}) {
	*min = append(*min, x.(int))
}

func (min *MinHeap) Pop() interface{} {
	old := *min
	n := len(old)
	x := old[n-1]
	*min = old[0 : n-1]
	return x
}

type MaxHeap []int

func (max MaxHeap) Len() int           { return len(max) }
func (max MaxHeap) Less(i, j int) bool { return max[i] > max[j] }
func (max MaxHeap) Swap(i, j int)      { max[i], max[j] = max[j], max[i] }

func (max *MaxHeap) Push(x interface{}) {
	*max = append(*max, x.(int))
}

func (max *MaxHeap) Pop() interface{} {
	old := *max
	n := len(old)
	x := old[n-1]
	*max = old[0 : n-1]
	return x
}

func main() {
	var n, d int
	fmt.Fscanf(stdin, "%d %d", &n, &d)
	upper := &MinHeap{}
	heap.Init(upper)
	lower := &MaxHeap{}
	heap.Init(lower)
	previous := make([]int, 0, d)
	notifications := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)

		var median float64
		if i >= d {
			if lower.Len() > upper.Len() {
				v := (*lower)[0]
				median = float64(v)
			} else {
				a := (*lower)[0]
				b := (*upper)[0]
				median = float64(a+b) / 2.0
			}
			if float64(x) >= 2*median {
				notifications++
			}
			toRemove := previous[0]
			previous = previous[1:]
			if toRemove <= (*lower)[0] {
				for i, v := range *lower {
					if toRemove == v {
						heap.Remove(lower, i)
						break
					}
				}
			} else {
				for i, v := range *upper {
					if toRemove == v {
						heap.Remove(upper, i)
						break
					}
				}
			}
		}
		if lower.Len() == 0 || x <= (*lower)[0] {
			heap.Push(lower, x)
		} else {
			heap.Push(upper, x)
		}
		for upper.Len() > lower.Len() {
			tmp := heap.Pop(upper)
			heap.Push(lower, tmp)
		}
		for lower.Len() > upper.Len()+1 {
			tmp := heap.Pop(lower)
			heap.Push(upper, tmp)
		}
		previous = append(previous, x)
	}
	fmt.Fprintln(stdout, notifications)
}
