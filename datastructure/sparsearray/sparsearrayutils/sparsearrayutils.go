package sparsearrayutils

import (
	"fmt"
)

// SparseToArray 稀疏数组->原始数组
func SparseToArray(sparseArray [][]int) (orangeArray [][]int) {
	row := sparseArray[0][0]          // 行数
	col := sparseArray[0][1]          // 列数
	defunctValue := sparseArray[0][2] // 默认值

	// 循环创立默认值的二维切片
	for i := 0; i < row; i++ {
		orangeArray = append(orangeArray, make([]int, col))
		for j := 0; j < col; j++ {
			orangeArray[i][j] = defunctValue
		}
	}

	// 循环填写非默认值数据
	for i := 1; i < len(sparseArray); i++ { // 稀疏切片跳过第一行
		orangeArray[sparseArray[i][0]][sparseArray[i][1]] = sparseArray[i][2] // 稀疏切片第0列是行号，第1列是列号，第2列是值，
	}

	return
}

// ArrayToSparse 原始数组->稀疏数组
func ArrayToSparse(orangeArray [][]int, defunctValue int) (sparseArray [][]int) {
	row := len(orangeArray)                                          // 行数
	col := len(orangeArray[0])                                       // 列数
	sparseArray = append(sparseArray, []int{row, col, defunctValue}) // 定义稀疏切片的首行
	for i, v1 := range orangeArray {
		for j, v2 := range v1 {
			if v2 != defunctValue {
				sparseArray = append(sparseArray, []int{i, j, v2}) // 稀疏切片加一行
			}
		}
	}
	return
}

// PrintSlice 打印二维切片
func PrintSlice(sli [][]int) {
	row := len(sli[0][:])
	fmt.Printf("行列\t")
	for i := 0; i < row; i++ {
		fmt.Printf("\t%v", i)
	}
	fmt.Println()
	//fmt.Println("      行\t      列\t      值")
	for i, v1 := range sli {
		fmt.Printf("%v\t", i)
		for _, v2 := range v1 {
			fmt.Printf("\t%v", v2)
		}
		fmt.Println()
	}
}

//
//// ArrayToSlice 二维数组->二维切片
//func ArrayToSlice(arr interface{}) (retSlice [][]int) {
//	arra := arr.([][]int)
//	fmt.Println("      行\t      列\t      值")
//	for i, v1 := range arr {
//		for j, v2 := range v1 {
//			fmt.Printf("%v\t      %v列\t      %v值\n", i, j, v2)
//		}
//	}
//}
