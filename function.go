// Package algorithms package demonstrates some basic algorithms
// Binary Sort , Select Sort , Recursive , Quick Sort , Hash table
package algorithms

import (
	"fmt"
	"math/rand"
)

// MakeArray creating an empty array, size , return x.
func MakeArray(size int) []int {
	x := make([]int, size)
	return x
}

// RandomFillArray filling an array with random numbers
func RandomFillArray(arr []int, rangeRandom int) []int {

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(rangeRandom)
	}
	return arr
}

// BinarySearch sorting array, arr is input array,
// item search value , return arr.
func BinarySearch(arr []int, item int) int {

	low := 0
	high := len(arr) - 1

	SelectorSort(arr)

	for low <= high {

		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// SelectorSort sorting array, arr is input array.
// return arr
func SelectorSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		var tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}

// Recursive search
func Recursive(i int) int {
	fmt.Println(i)
	if i == 0 {
		return 10
	} else {
		Recursive(i - 1)
	}
	return 0
}

// Fact calc
func Fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * Fact(x-1)
	}
}

// QuickSort is fast sorting.
func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	//fmt.Println(pivot)
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}

// MinMax is finds the minimum and maximum elements in the array.
func MinMax(array []float64) (min float64, max float64) {
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
