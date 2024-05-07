package sorts

import (
	"math"
)

// BubbleSort 冒泡排序
func BubbleSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// arr[j], arr[j+1] = arr[j+1], arr[j]
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				// 表示有数据交换
				flag = true
			}
		}
		if !flag {
			break // 没有数据交换，推出循环
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j] // 数据移动
			} else {
				break
			}
		}
		arr[j+1] = value // 插入数据
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

/*

归并排序（Merge Sort）：
优点：
1. 稳定性：归并排序是一种稳定的排序算法，相等元素的相对位置不会改变。
2. 高效性：归并排序的时间复杂度为 O(n log n)，在所有情况下都保持稳定。
3. 适合大规模数据：由于归并排序是一种分治算法，它可以很好地利用递归与迭代来处理大规模数据，适用于外部排序。
4. 对链表排序效率高：由于归并排序只需要对链表进行节点指针的重新连接，因此在链表排序中更加高效。

缺点：
1. 需要额外空间：归并排序需要额外的空间来存储临时数组，空间复杂度为 O(n)。
2. 不是原地排序：归并排序通常需要额外的空间来存储临时数组，因此不适用于内存受限的情况。

适用场景：
1. 适用于对大规模数据进行排序，特别是外部排序。
2. 当需要稳定排序时，归并排序是一个不错的选择。
3. 对链表进行排序时，归并排序效率较高。
*/

// TODO 递归有点难度，没咋看懂
// MergeSort 归并排序
func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = arr[i]
		i++
		k++
	}
	for j <= end {
		tmp[k] = arr[j]
		j++
		k++
	}
	copy(arr[start:end+1], tmp)

}

/*
快速排序（Quick Sort）：
优点：
1. 原地排序：快速排序是一种原地排序算法，不需要额外的空间来存储临时数组，空间复杂度为 O(nlogn)。
2. 高效性：在平均情况下，快速排序的时间复杂度为 O(nlogn)，在实践中性能非常好。
3. 算法简单：快速排序的实现相对简单，容易理解和实现。

缺点：
1.对于特定数据集，可能会出现最坏情况：在某些情况下（例如已经有序的数组），快速排序可能会退化为 O(n^2) 的时间复杂度。
2. 不稳定：快速排序是一种不稳定的排序算法，相等元素的相对位置可能会改变。

适用场景：
1. 适用于需要原地排序且性能要求高的场景，例如大部分编程语言内置的排序算法就是快速排序。
2. 当需要简单高效的排序算法时，快速排序是一个不错的选择。
3. 对于大部分情况下的排序需求，快速排序都能够较好地满足。


综合比较：
如果你需要稳定的排序算法，并且对额外空间的使用不敏感，那么归并排序可能更适合你。
如果你对内存占用敏感，并且不需要保持相等元素的相对位置不变，那么快速排序可能更适合你。
在实际应用中，对数据特征的了解以及算法的实现都会影响选择。

*/

// QuickSort 快速排序
func QuickSort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i

}

// TODO 没学会 线性排序
// BucketSort 桶排序
func BucketSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	// 获取待排序数组中最大值
	max := getMax(arr, n)
	buckets := make([][]int, n)
	index := 0

	for i := 0; i < n; i++ {
		index = arr[i] * (n - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < n; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i], bucketLen) // 桶内做快速排序
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

func getMax(arr []int, length int) int {
	max := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max

}

// 计数排序（Counting sort）
func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}
	var max int = math.MinInt32
	// 循环获取最大值
	for i := range a {
		if a[i] > max {
			max = a[i]
		}

	}

	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}

// 获取数组中的最大值
func getMax1(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// 对数组按照指定位数进行计数排序
func countSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) // 用于存储排序结果的临时数组
	count := make([]int, 10) // 计数数组，用于记录数字出现的次数

	// 初始化计数数组
	for i := 0; i < 10; i++ {
		count[i] = 0
	}

	// 统计每个数字出现的次数
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// 将 count[i] 转换为数字在输出数组中的实际位置
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 根据计数数组中的信息，将数字放入输出数组中
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	// 将排序结果复制回原数组
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// 基数排序函数
func RadixSort(arr []int) {
	max := getMax1(arr) // 获取数组中的最大值
	// 从个位开始，对数组进行计数排序，依次按照个位、十位、百位...进行排序
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}
