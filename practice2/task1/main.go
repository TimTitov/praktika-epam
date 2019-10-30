package main

import (
	"fmt"
	c "practice2/task1/collection"
)

func main() {
	collection := c.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(i)
	}
	collection.Print()
	fmt.Println("Get 3 = ", collection.Get(3))
	collection.Remove(4)
	collection.Print()
	fmt.Println("Get 5 = ", collection.Get(5))
	fmt.Println("Get First = ", collection.First())
	fmt.Println("Get Last = ", collection.Last())
	fmt.Println("Get Length = ", collection.LengthCheck())
}
