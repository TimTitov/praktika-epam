package funcarr

import "fmt"

// ChangeMass change array.
func ChangeMass(item string, arr [][]string) {
	for i := 0; i < len(arr); i++ {
		arr[i][i] = item
		arr[i][len(arr)-1-i] = item
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
func PrintMass(arr [][]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
