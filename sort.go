package main

import "fmt"

// caibo
// 手撕的排序算法
// 选择排序、快速排序、希尔排序、堆排序不是稳定的排序算法，而冒泡排序、插入排序、归并排序和基数排序是稳定的排序算法。

// QuickSort 快速排序
func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	x := arr[i]
	for i < j {
		for i < j && x <= arr[j] {
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}
		for i < j && arr[i] <= x {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}
	}
	arr[i] = x
	QuickSort(arr, l, i-1)
	QuickSort(arr, i+1, r)
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for i := 0; i < length; i++ {
		var flag bool
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// InsertSort 插入排序
func InsertSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for i := 1; i < length; i++ {
		j := i - 1
		x := arr[i]
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}

// SelectSort 简单选择排序
func SelectSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

// HeapSort 堆排序
func HeapSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, length)
	}
	for j := length - 1; j > 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		adjustHeap(arr, 0, j)
	}
}

// adjustHeap 堆排序
func adjustHeap(arr []int, i int, length int) {
	temp := arr[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > temp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}

// MergeSort 归并排序
func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		MergeSort(arr[:mid])
		MergeSort(arr[mid:])
		merge(arr, mid)
	}
}

// merge 归并排序
func merge(arr []int, mid int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	l, r := 0, mid
	cur := 0
	for l < mid && r < len(temp) {
		if temp[l] <= temp[r] {
			arr[cur] = temp[l]
			l++
		} else {
			arr[cur] = temp[r]
			r++
		}
		cur++
	}
	for l < mid {
		arr[cur] = temp[l]
		l++
		cur++
	}
	for r < len(temp) {
		arr[cur] = temp[r]
		r++
		cur++
	}
}

// BucketSort 桶排序
func BucketSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	tmp := make([]int, 100001)
	for i := 0; i < length; i++ {
		tmp[arr[i]+50000]++
	}
	cur := 0
	for i := 0; i < len(tmp); i++ {
		for tmp[i] > 0 {
			arr[cur] = i - 50000
			cur++
			tmp[i]--
		}
	}
}

// RadixSort 基数排序(正负数不能一起)
func RadixSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	d := 6
	tmp := make([]int, length)
	radix := 1
	for i := 1; i <= d; i++ {
		cnt := make([]int, 10)
		for j := 0; j < length; j++ {
			k := (arr[j] / radix) % 10
			fmt.Println(k)
			cnt[k]++
		}
		for j := 1; j < 10; j++ {
			cnt[j] = cnt[j] + cnt[j-1]
		}
		// 这里从后面开始很重要 升序
		for j := length - 1; j > 0; j-- {
			k := (arr[j] / radix) % 10
			tmp[cnt[k]-1] = arr[j]
			cnt[k]--
		}
		copy(arr, tmp)
		radix *= 10
	}
}
