package main

import (
	"fmt"
	"os"
	f "practice1/funcarr"
)

func main() {
	var size int
	var item string
	var diagitem string

	fmt.Print("Введите размер массива: ")
	fmt.Fscan(os.Stdin, &size)
	fmt.Print("Введите символ заполнения: ")
	fmt.Fscan(os.Stdin, &item)
	fmt.Print("Введите символ диагонали: ")
	fmt.Fscan(os.Stdin, &diagitem)

	arr := f.CreateMass(item, size)
	f.ChangeMass(diagitem, arr)
	f.PrintMass(arr)
}
