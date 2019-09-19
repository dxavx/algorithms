// Binary search
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// ----
	var x = MakeArray(10) 				// new array of 10 zeros
	fmt.Println(x) 							// viewing an empty array
	//fmt.Println(&x[0]) 					// address of the first element
	RandonFillArray(x, 100) 	// filling an array with random numbers
	fmt.Println(x) 							// viewing an empty array
	sort.Ints(x)							// sorting array
	fmt.Println(x)							// viewing an sorting array
	fmt.Println(BinarySearch(x,50)) 	// search "50" in array
}

											// creating an empty array
func MakeArray (size int) ([]int) {
	x := make([]int,size)
	return x
}
											// illing an array with random numbers
func RandonFillArray (arr []int , range_random int) []int {

	//fmt.Println(&arr[0])
	rand.Seed(time.Now().Unix())
	for i := 0 ; i < len(arr); i++  {
		arr[i] = rand.Intn(range_random)
	}
	return arr
}

func BinarySearch (arr []int, item int) int {

	//fmt.Println(&arr[0])
	low := 0 								// lower bound of range
	high := len(arr) - 1 					// height bound of range

	for low <= high {
		mid := (low + high) / 2				// reduce the search range
		guess := arr[mid]
		if guess == item { return mid }
		if guess > item { high = mid - 1 } else { low = mid + 1 }
		fmt.Println(low, high)
	}
	return -1
}

func Smallest (arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i,  _ := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
		return smallest_index
	}
	return smallest_index
}
