package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

//var tran = make([][]int, 100, 100)
var tran [][]int


func max_path_of_triangle(root int) int {
	//var dp = make([][]int, 100, 100)
	var dp [][]int
	dp = tran
	for i := root-2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = mmax(dp[i+1][j], dp[i+1][j+1]) + tran[i][j]
		}
	}

	return dp[0][0]
}


func load_triangle_data(file string) int{
	f, err := os.Open(file)
	if err != nil {
		return 0
	}
	buf := bufio.NewReader(f)
	var i = 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handle(line, i+1)
		i++
		if err != nil {
			if err == io.EOF {
				return i
			}
			return 0
		}
	}


	return i
}

func handle(line string, deep int)  {
	var ss = make([]string, 100)
	ss = strings.Fields(line)
	var temp []int
	for _, node := range ss{
		tt , _ := strconv.Atoi(node)
		temp = append(temp, tt)
	//	fmt.Println(temp)
	}
	tran = append(tran, temp)
}

func mmax(a , b int) int {
	if a >= b {
		return a
	}

	return b
}

func main()  {
	var file = "file"
	root := load_triangle_data(file)
	if root == 0 {
		fmt.Println("wrong")
	}
	ans := max_path_of_triangle(root)
	fmt.Println(ans)

	}


//
//int max_path_of_triangle(int **root)
//{
//return 0;
//}
//
//// load triangle data, return NULL if error
//int** load_triangle_data(char *file_name)
//{
//return NULL;
//}
//
//int main(void)
//{
//printf("hello, world\n");
//}
