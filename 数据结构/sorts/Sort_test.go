package sorts

import (
	"fmt"
	"math/rand"
	"testing"
)

// TestBubbleSort 冒泡排序
func TestBubbleSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	BubbleSort(arr, len(arr))
	t.Log(arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 11, 10}
	fmt.Println("排序前：", arr)
	InsertionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	SelectionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestMergeSort(t *testing.T) {
	// arr := []int{5, 4}
	// MergeSort(arr, len(arr))
	// t.Log(arr)

	arr := createRandomArr(100)
	t.Logf("排序前：%v", arr)
	MergeSort(arr, len(arr))
	t.Logf("排序后：%v", arr)
}

func createRandomArr(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(10000000)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	// arr := []int{5, 4}
	// QuickSort(arr, len(arr))
	// t.Log(arr)

	arr := createRandomArr(100)
	t.Logf("排序前：%v", arr)
	QuickSort(arr, len(arr))
	t.Logf("排序后：%v", arr)
}

func TestBucketSort(t *testing.T) {

	arr := createRandomArr(100)
	BucketSort(arr, len(arr))
	t.Log(arr)
}

func TestRadixSort(t *testing.T) {
	arr := createRandomArr(100)

	RadixSort(arr)
	t.Logf("排序后：%v", arr)
}
