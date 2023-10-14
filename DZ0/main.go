package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func countingSort(arr []int) []int {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	count := make([]int, max-min+1)
	for _, v := range arr {
		count[v-min]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(arr))
	for _, v := range arr {
		result[count[v-min]-1] = v
		count[v-min]--
	}

	return result
}

func createArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

func createRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}

	return arr
}

func measureTime(start time.Time, nameSort string) {
	elapsed := time.Since(start)
	fmt.Printf("Время сортировки %s: %d ticks\n", nameSort, elapsed.Nanoseconds())
}

func main() {
	var flag, size int
	var arr []int

	fmt.Print("Каким способом создаем массив? (1 - рандом, 2 - руками): ")
	fmt.Scanln(&flag)

	if flag == 1 {
		fmt.Print("Введите размер массива: ")
		fmt.Scanln(&size)
		arr = createRandomArray(size)
	} else if flag == 2 {
		fmt.Print("Введите размер массива: ")
		fmt.Scanln(&size)
		arr = createRandomArray(size)
	} else {
		fmt.Print("Ошибка ввода")
		return
	}

	fmt.Println("Изначальный массив: ", arr)

	startBubble := time.Now()
	bubbleSort(arr)
	measureTime(startBubble, "Bubble")

	startInsertion := time.Now()
	bubbleSort(arr)
	measureTime(startInsertion, "Insertion")

	startMerge := time.Now()
	bubbleSort(arr)
	measureTime(startMerge, "Merge")

	startCounting := time.Now()
	bubbleSort(arr)
	measureTime(startCounting, "Counting")
}
