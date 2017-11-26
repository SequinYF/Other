package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

type va struct {
	Sort
}

type Sort interface {
	Qsort(numbers []int, left, right int)
	Merge(numbers []int, left, right int)
}

func (v *va)Qsort(numbers []int, left, right int) {
	if left < right {
		q := Partition(numbers, left, right)
		go	v.Qsort(numbers, left, q - 1)
		go	v.Qsort(numbers, q + 1, right)
	}
}

func Partition(numbers []int, left, right int) int{
	i := left - 1
	for j := left; j <= right - 1; j++ {
		if numbers[j] < numbers[right] {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[right] = numbers[right], numbers[i+1]
	return i + 1
}

func (v *va)Merge(numbers []int, left, right int) {

}


func main() {

	r := bufio.NewReader(os.Stdin)

	for {
		in, _, _:= r.ReadLine()
		line := string(in)
		number := strings.Split(line, " ")

		numbers := make([]int, len(number))

		for i, num := range number {
			numbers[i], _ = strconv.Atoi(num)
		}

		numbers1 := numbers
		v := &va{}
		v.Qsort(numbers, 0, len(numbers) -1)
		fmt.Println("qsort")
		for _, num := range numbers {
			fmt.Print(num ," ")
		}
		v.Merge(numbers1, 0, len(numbers1) - 1)
		fmt.Println("merge")
		for _, num := range numbers1 {
			fmt.Print(num ," ")
		}
	}
}
