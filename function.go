package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

// MakeArray creating an empty array
func MakeArray(size int) []int {
	x := make([]int, size)
	return x
}

// RandomFillArray filling an array with random numbers
func RandomFillArray(arr []int, rangeRandom int) []int {

	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(rangeRandom)
	}
	return arr
}

// BinarySearch
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

// SelectorSort
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
func Recursive(i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	} else {
		Recursive(i - 1)
	}
}

// Fact calc
func Fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * Fact(x-1)
	}
}

// QuickSort , fast sorting
func QuickSort(arr []int) []int {

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

func Example() {

	// ---- Select Sort
	var x1 = MakeArray(20)
	RandomFillArray(x1, 1000)
	fmt.Println(x1)
	fmt.Println(SelectorSort(x1))

	// ------ Recursive
	Recursive(30)

	//------- Fact
	fmt.Println(Fact(7))

	//------- Quick Sort
	var x2 = MakeArray(10)
	RandomFillArray(x2, 1000)
	fmt.Println(x2)
	fmt.Println(QuickSort(x2))

	//----- Hash -------------
	m := make(map[string]int)
	m1 := new(map[string]int)
	var m2 map[string]int
	m3 := map[string]int{"apple": 99, "orange": 100}

	// Insert data
	m3["carrot"] = 101

	// Delete data
	delete(m3, "apple")

	// Search data
	xx, aa := m3["orange"]
	fmt.Println("target value", xx, aa)

	fmt.Println(m, m1, m2, m3)
	//-------------------------

	// ---- Graph ---------------
	graph := make(map[string][]string)
	c := make(chan string, 300)

	graph["you"] = []string{"video", "audio", "tv"}
	graph["pat"] = []string{"js", "php", "nodajs"}
	graph["danya"] = []string{"qa", "test", "graph"}
	graph["max"] = []string{}

	fmt.Println("Test Graph", graph)
	// Pipe
	c <- graph["you"][0]
	fmt.Println(<-c)

	//---------------------------

}
