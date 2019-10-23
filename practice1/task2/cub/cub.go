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
	countarr := make([]int, 12)
	for i := 0; i < size; i++ {

		switch arr[i] {
		case 2:
			countarr[0]++
		case 3:
			countarr[1]++
		case 4:
			countarr[2]++
		case 5:
			countarr[3]++
		case 6:
			countarr[4]++
		case 7:
			countarr[5]++
		case 8:
			countarr[6]++
		case 9:
			countarr[7]++
		case 10:
			countarr[8]++
		case 11:
			countarr[9]++
		case 12:
			countarr[10]++
		}

	}
	for j := range countarr {
		fmt.Print(j + 2)
		fmt.Print("---")
		fmt.Printf("%.1f \n", (100 / float64(countarr[j])))

		//fmt.Printf(" %v", inc(100/countarr[j]))
	}

	//var xx1 float64 = float64(100 / float64(x2))
	//x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 = 100/x2, 100/x3, 100/x4, 100/x5, 100/x6, 100/x7, 100/x8, 100/x9, 100/x10, 100/x11, 100/x12
	//	fmt.Printf("%.1f", (100 / x2))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// func procent(){

// }
