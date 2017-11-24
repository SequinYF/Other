package main

import (
	"flag"
	"fmt"

	"os"
	"bufio"
	"io"
	"strconv"
	//文件读写和字符串处理
	"time"
	"sorter/sorter/algorithms/qsort"
	"sorter/sorter/algorithms/bubblesort"
)


//定义命令行中的参数
//name:名称 value：默认值 usage：解释说明
var infile *string = flag.String("i", "unsorted.dat", "File contains values" +
	" for sorting")
var outfile *string = flag.String("o", "sorted.dat"	, "File to receive " +
	"storted values")
var algorithm *string  = flag.String("a", "qsort", "Sort algorithm")


func readValues(infile string) (values []int, err error) {
	//func Open(name string) (file *File, err error)
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Faild to open the input file ", infile)
		return
	}

	defer file.Close()

	//将 br 封装成一个带缓存的 bufio.Reader 对象, size:4096
	//func NewReader(br io.Reader) *Reader
	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		//(b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		//bytes -> string
		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		//values := make([]int, 0)
		values = append(values, value)
	}

	return
}


func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Faild to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {


	//解析参数
	flag.Parse()

	//flag.String()返回的是参数的内存地址，通过加*获取值
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("The sorting algoritm", *algorithm, "is either unkonwn or unsupported.")
		}

		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

