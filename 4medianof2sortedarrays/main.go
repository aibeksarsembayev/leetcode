package main

import "fmt"

func main() {

	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge 2 arrays
	mergedArray := getMergedArray(nums1, nums2)
	fmt.Println(mergedArray)
	// sort merged array
	sortedMergedArray := sortMergedArray(mergedArray)
	fmt.Println(sortedMergedArray)
	// find median
	median := getMedian(sortedMergedArray)
	fmt.Println(median)
	return median
}

func getMedian(arr []int) float64 {
	var median float64
	if len(arr)%2 == 0 {
		median = (float64(arr[(len(arr)-1)/2]) + float64(arr[(len(arr)-1)/2+1])) / 2
	} else {
		median = float64(arr[(len(arr)-1)/2])
	}
	return median
}

func sortMergedArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func getMergedArray(n1, n2 []int) []int {
	var mergedArray []int
	mergedArray = n1
	for j := 0; j < len(n2); j++ {
		mergedArray = append(mergedArray, n2[j])
	}
	return mergedArray
}

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n))
