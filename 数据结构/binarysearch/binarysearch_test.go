package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	if BinarySearch(a, 8) != 4 {
		t.Fatal(BinarySearch(a, 3))
	}
	if BinarySearch(a, 4) != -1 {
		t.Fatal(BinarySearch(a, 4))
	}
}

func TestBinarySearchRecursive(t *testing.T) {

	a := []int{1, 3, 5, 6, 8}
	index := BinarySearchRecursive(a, 2)
	if index == -1 {
		t.Fatal("not found")
	}
	t.Logf("index: %d", index)
}

func TestBinarySearchFirst(t *testing.T) {

	a := []int{1, 2, 2, 2, 3, 4}
	index := BinarySearchFirst(a, 2)
	if index == -1 {
		t.Fatal("not found")
	}
	t.Logf("index: %d", index)

}
