package main

import (
	"fmt"
	"os"
	c "practice1-2/cub"
)

func main() {
	var size int
	fmt.Print("Введите количество бросков: ")
	fmt.Fscan(os.Stdin, &size)
	var arr []int
	c.Core(size, arr)
}
