package binarysearch

func BinarySearch(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1

	for low <= high {
		// 计算中间数值
		mid := (low + high) / 2
		if a[mid] == v {
			return mid
		} else if v < a[mid] {
			//如果目标值 v 小于中间元素，则将搜索范围缩小为左半部分，即将 high 更新为 mid - 1
			high = mid - 1
		} else {
			// 如果目标值 v 大于中间元素，则将搜索范围缩小为右半部分，即将 low 更新为 mid + 1。
			low = mid + 1
		}
	}
	return -1
}

// 递归搜索
func BinarySearchRecursive(a []int, v int) int {
	l := len(a)
	return bsearchInternally(a, 0, l-1, v)
}

func bsearchInternally(a []int, low, high, v int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if a[mid] == v {
		return mid
	} else if v > a[mid] {
		return bsearchInternally(a, mid+1, high, v)
	} else {
		return bsearchInternally(a, low, mid-1, v)
	}
}

// 查找第一个值等于给定值的元素
func BinarySearchFirst(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] < v {
			low = mid + 1
		} else if a[mid] > v {
			high = mid - 1
		} else {
			if mid == 0 || a[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func BinarySearchLast(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] < v {
			low = mid + 1
		} else if a[mid] > v {
			high = mid - 1
		} else {
			if mid == n-1 || a[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGt(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func BinarySearchLastLt(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] <= v {
			if mid == n-1 || a[mid+1] > v {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
