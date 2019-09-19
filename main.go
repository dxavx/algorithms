// Binary search
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// ---- Binary Sort
	var x = MakeArray(10) 				// new array of 10 zeros
	fmt.Println(x) 							// viewing an empty array
	RandonFillArray(x, 100) 	// filling an array with random numbers
	fmt.Println(x) 							// viewing an empty array
	sort.Ints(x)							// sorting array
	fmt.Println(x)							// viewing an sorting array
	fmt.Println(BinarySearch(x,50)) 	// search "50" in array

	// ---- Select Sort
	var x1 = MakeArray(20)
	RandonFillArray(x1,1000)
	fmt.Println(x1)
	fmt.Println(SelectorSort(x1))
}

											// creating an empty array
func MakeArray (size int) ([]int) {
	x := make([]int,size)
	return x
}
											// illing an array with random numbers
func RandonFillArray (arr []int , range_random int) []int {

	rand.Seed(time.Now().Unix())
	for i := 0 ; i < len(arr); i++  {
		arr[i] = rand.Intn(range_random)
	}
	return arr
}

func BinarySearch (arr []int, item int) int {

	low := 0 								// lower bound of range
	high := len(arr) - 1 					// height bound of range

	for low <= high {
		mid := (low + high) / 2				// reduce the search range
		guess := arr[mid]
		if guess == item { return mid }
		if guess > item { high = mid - 1 } else { low = mid + 1 }
	}
	return -1
}
											// Selection sort
func SelectorSort (arr []int) []int {

	for i := 0 ; i < len(arr) -1 ; i++ {
		var min  = i
		for j := i + 1 ; j < len(arr); j++ {
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

