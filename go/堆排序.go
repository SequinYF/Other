package main

import (
	"fmt"
)

func swap(a, b *int64)  {
	var tep int64
	tep = *a
	*a = *b
	*b = tep
}

func makemaxheap(arr []int64, start, end int)  {

	index := start
	ls := index << 1 + 1

	for ; ls <= end; {

		if ls < end && arr[ls] < arr[ls + 1] {
			ls++
		}
		if arr[index] < arr[ls] {
			swap(&arr[index], &arr[ls])
		} else {
			break
		}

		index = ls
		ls = 2 * ls + 1

	}
}



func makeminheap(arr []int64, start, end int)  {
	index := start
	ls := index << 1 + 1

	for ; ls <= end; {

		if ls < end && arr[ls] > arr[ls + 1] {
			ls++
		}
		if arr[index] > arr[ls] {
			swap(&arr[index], &arr[ls])
		} else {
			break
		}

		index = ls
		ls = 2 * ls + 1

	}
}

func hsortdesc(arr []int64, n int)  {

	for i := n/2-1; i >= 0; i-- {
		makeminheap(arr, i, n-1)
	}
	for i := n-1; i > 0; i-- {
		swap(&arr[i], &arr[0])
		makeminheap(arr, 0, i-1)
	}
}

func hsortasc(arr []int64, n int)  {
	for i:= n/2-1; i >= 0; i-- {
		makemaxheap(arr, i, n-1)
	}

	for i := n-1; i > 0; i-- {
		swap(&arr[i], &arr[0])
		makemaxheap(arr, 0, i-1)
	}
}
func main() {

	var arr = []int64{20,30,90,40,70,110,60,10,100,50,80}
	//hsortasc(arr, 11)
	hsortdesc(arr, 11)

	for _, node := range arr {
		fmt.Println(node)
	}

}
