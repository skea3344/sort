package main

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


// 归并排序 
