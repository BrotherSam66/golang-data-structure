package main

import (
	"sparsearray/sparsearrayutils"
)

func main() {
	arr := [8][5]int{
		{0, 0, 7, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 2, 0, 0},
		{9, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 4, 303, -5, 0},
	}

	// 二维数组->二维切片
	var arrSlice [][]int
	for i := 0; i < len(arr); i++ {
		arrSlice = append(arrSlice, arr[i][:]) // 单行切片追加到总切片里
	}

	sparsearrayutils.PrintSlice(arrSlice)
	spa := sparsearrayutils.ArrayToSparse(arrSlice, 0) // 标准数组->稀疏数组。默认值是0
	sparsearrayutils.PrintSlice(spa)
	newArr := sparsearrayutils.SparseToArray(spa) // 稀疏数组->标准数组
	sparsearrayutils.PrintSlice(newArr)
}
