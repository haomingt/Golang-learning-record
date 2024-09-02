package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("请输入分数：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
