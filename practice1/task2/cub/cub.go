package cub

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func Core(size int, arr []int) {
	var x int = 0
	for i := 0; i < size; i++ {
		arr = append(arr, randomInt(1, 7)+randomInt(1, 7))
		x++
	}

	//var x2,x3,x4,x5,x6,x7,x8,x9,x10,x11,x12 int
	x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	//sort.Ints(arr)
	//fmt.Println(arr)
	for i := 0; i < size; i++ {

		switch arr[i] {
		case 2:
			x2++
		case 3:
			x3++
		case 4:
			x4++
		case 5:
			x5++
		case 6:
			x6++
		case 7:
			x7++
		case 8:
			x8++
		case 9:
			x9++
		case 10:
			x10++
		case 11:
			x11++
		case 12:
			x12++
		}

	}
	fmt.Println(float64(x2))
	//var xx1 float64 = float64(100 / float64(x2))
	//x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 = 100/x2, 100/x3, 100/x4, 100/x5, 100/x6, 100/x7, 100/x8, 100/x9, 100/x10, 100/x11, 100/x12
	fmt.Printf("%.1f", (100 / x2))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// func procent(){

// }
