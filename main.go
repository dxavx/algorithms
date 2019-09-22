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
	RandomFillArray(x, 100) 	// filling an array with random numbers
	fmt.Println(x) 							// viewing an empty array
	sort.Ints(x)							// sorting array
	fmt.Println(x)							// viewing an sorting array
	fmt.Println(BinarySearch(x,50)) 	// search index "50" in array

	// ---- Select Sort
	var x1 = MakeArray(20)
	RandomFillArray(x1,1000)
	fmt.Println(x1)
	fmt.Println(SelectorSort(x1))

	// ------ Recursive
	Recursive(30)

	//------- Fact
	fmt.Println(Fact(7))

	//------- Quick Sort
	var x2 = MakeArray(10)
	RandomFillArray(x2,1000)
	fmt.Println(x2)
	fmt.Println(QuickSort(x2))
}

											// creating an empty array
func MakeArray (size int) []int {
	x := make([]int,size)
	return x
}
											// Filling an array with random numbers
func RandomFillArray (arr []int , rangeRandom int) []int {

	rand.Seed(time.Now().Unix())
	for i := 0 ; i < len(arr); i++  {
		arr[i] = rand.Intn(rangeRandom)
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

func Recursive ( i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	} else {
		Recursive(i-1)
	}
}

func Fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return  x * Fact(x-1)
	}
}

func QuickSort (arr []int) []int{

	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	fmt.Println(pivot)
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

