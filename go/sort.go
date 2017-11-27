package main

import (
   // "bufio"
   // "os"
   // "strings"
    "fmt"
   // "strconv"
    "time"
    "math/rand"
)

type va struct {
    Sort
}

type Sort interface {
    Qsort(numbers []int, left, right int)
    Merge_sort(numbers []int, left, right int, temp []int)
}

func (v *va)Qsort(numbers []int, left, right int) {
    if left < right {
        q := Partition(numbers, left, right)
        v.Qsort(numbers, left, q - 1)
        v.Qsort(numbers, q + 1, right)
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

/*
func merge(numbers []int, left, mid, right int, temp []int) {
    i, j, t := left, mid + 1, 0
    for i < mid && j <=right {
        if numbers[i] <= numbers[j] {
            temp[t] = numbers[i]
            t++
            i++
        } else {
            temp[t] = numbers[j]
            t++
            j++
        }
    }

    for i <= mid {
        temp[t] = numbers[i]
        i++
        t++
    }

    for j <= right {
        temp[t] = numbers[j]
        t++
        j++
    }

    t = 0

    for left <= right {
        numbers[left] = temp[t]
        t++
        left++
    }
}

*/
/*

func (v *va)Merge_sort(numbers []int, left, right int, temp []int) {

    if left < right {
        mid := (left + right) / 2
        go v.Merge_sort(numbers, left, mid, temp)
        go v.Merge_sort(numbers, mid + 1, right, temp)
        merge(numbers, left, mid, right, temp)
    }
}
*/

func main() {

   // r := bufio.NewReader(os.Stdin)

    //for {
        /*
        in, _, _:= r.ReadLine()
        line := string(in)
        number := strings.Split(line, " ")

        numbers := make([]int, len(number))

        for i, num := range number {
            numbers[i], _ = strconv.Atoi(num)
        }
        */

        numbers := make([]int, 10000)

        for i := 0; i < 100; i++ {
            numbers[i] = rand.Intn(100)
        }

 //       numbers1 := numbers
        v := &va{}
        t1 := time.Now()
        v.Qsort(numbers, 0, len(numbers) -1)
        t2 := time.Now()
        fmt.Println("The qsort sorting process costs", t2.Sub(t1), "to complete")
       /*
        for _, num := range numbers {
            fmt.Print(num ," ")
        }
        */
        /*
        temp := make([]int, len(numbers) + 2)
        t1 = time.Now()
        v.Merge_sort(numbers1, 0, len(numbers1) - 1, temp)
        t2 = time.Now()
        fmt.Println("\nThe merge sorting process costs", t2.Sub(t1), "to complete")
        */
        /*
        for _, num := range numbers1 {
            fmt.Print(num ," ")
        }
        */
    //}
}
