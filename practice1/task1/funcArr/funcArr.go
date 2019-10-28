package funcarr

import "fmt"

// ChangeMass change array.
func ChangeMass(item string, arr [][]string, size int) {
	var count = 0
	var count2 = size - 1
	for i := 0; i < size; i++ {
		arr[i][count] = item
		arr[i][count2] = item
		count++
		count2--
	}
}

// CreateMass creates array.
func CreateMass(item string, size int) [][]string {
	arr := make([][]string, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]string, size)
		for j := 0; j < size; j++ {
			arr[i][j] = item
		}
	}
	return arr
}

// PrintMass writes array to console.
func PrintMass(size int, arr [][]string) {
	for i := 0; i < size; i++ {
		fmt.Println(arr[i])
	}
}
