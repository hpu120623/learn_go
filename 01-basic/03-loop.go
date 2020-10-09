package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 读文件，输出内容
func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever()  {
	for {
		fmt.Println("hello go")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011-->1101
	)
	printFile("abc.txt")
	forever()
}
