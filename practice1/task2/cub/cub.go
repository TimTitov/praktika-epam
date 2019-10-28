package cub

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//Core collects and print statistics
func Core(size int, arr []int) {
	var x int = 0
	countarr := make([]int, 11)
	var value int
	for i := 0; i < size; i++ {
		value = randomInt(1, 7) + randomInt(1, 7)
		countarr[value-2]++
		arr = append(arr, value)
		x++
	}
	//fmt.Println("То что выпало: ", arr)
	//fmt.Println("Сколько выпало каждого: ", countarr)
	for i := 0; i < 11; i++ {
		fmt.Printf("%d - %2.1f%%\n", i+2, float64(countarr[i])*100.0/float64(size))
	}
}
