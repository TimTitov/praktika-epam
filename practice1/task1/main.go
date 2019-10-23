package main

import (
	f "practice1/funcArr"
)

func main() {
	size := 11
	arr := f.CreateMass("x", size)
	f.ChangeMass("-", arr, size)
	f.PrintMass(size, arr)
}
