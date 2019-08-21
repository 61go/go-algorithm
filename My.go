package main

import "fmt"

func main() {

	//var x = []int{3, 4, 1, 8, 7, 2323, 21, 2, 9}
	//sort(x)
	//for e, j := range x {
	//	println(e, j)
	//}
	var arr = []int{2, 4, 7, 18, 1, 8, 9, 12}
	merge(arr, 0, 4, len(arr)-1)
	for k, v := range arr {
		fmt.Printf("%d------%d\n", k, v)
	}
	//t1 := time.Now().Unix()
	//var ii = factor(2, 12)
	//println(ii)
	//t2 := time.Now().Unix()
	//println(t2 - t1)
}

/**
插入排序精要 插入排序是最简单的排序
*/
func sort(x []int) {
	for i := 1; i < len(x); i++ {
		for j := i; j > 0; j-- {
			if x[j-1] > x[j] {
				swap(x, j-1, j)
			}
		}
	}
}

/**
归并算法
*/
func merge(arr []int, leftPos int, rightPos int, rightEnd int) {
	var tmp = make([]int, len(arr))
	leftEnd := rightPos - 1
	var index = 0
	for ; leftPos <= leftEnd && rightPos <= rightEnd; index++ {
		if arr[leftPos] > arr[rightPos] {
			tmp[index] = arr[rightPos]
			rightPos++
		} else {
			tmp[index] = arr[leftPos]
			leftPos++
		}
	}
	for ; leftPos <= leftEnd; leftPos++ {
		tmp[index] = arr[leftPos]
	}

	for ; rightPos <= rightEnd; rightPos++ {
		tmp[index] = arr[rightPos]
		index++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = tmp[i]
	}
}

/**
怎么算乘方-0-0 乘方有效率的算法
*/
func factor(x int, n int) int {
	if n == 0 {
		return 1
	} else if n&1 == 1 {
		tt := factor(x, (n-1)/2)
		return tt * tt * x
	} else {
		tt := factor(x, n/2)
		return tt * tt
	}
}

/**
算法:交互数组两个位置上的数值 只限于整数
*/
func swap(arr []int, left int, right int) {
	arr[left] ^= arr[right]
	arr[right] ^= arr[left]
	arr[left] ^= arr[right]
}
