package main

import (
	"fmt"
	"math/rand"
	"sorting/utils"
	"time"
)

func main() {
	//sli := []int{1, 99, 109, 44, 90, 9, 56, 89, 45, 3, 79}

	//sli := []int{1, 99, 99, 44, 90, 9, 56, 89, 45, 3, 79}

	length := 10000
	startTime := time.Now().UnixMicro()
	stopTime := time.Now().UnixMicro()
	timeDuration := stopTime - startTime
	sli := make([]int, length)

	// 选择排序测试
	for i := 0; i < length; i++ {
		sli[i] = rand.Intn(1000000)
	}
	startTime = time.Now().UnixMicro()
	fmt.Println(sli)
	utils.SelectSort(&sli)
	fmt.Println(sli)
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("选择排序执行时间SelectSort，毫秒：", timeDuration)

	// 插入排序
	for i := 0; i < length; i++ {
		sli[i] = rand.Intn(1000000)
	}
	startTime = time.Now().UnixMicro()
	fmt.Println(sli)
	utils.InsertSort(&sli)
	fmt.Println(sli)
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("插入排序执行时间InsertSort，毫秒：", timeDuration)
	rand.Seed(time.Now().UnixNano())

	// 快速排序测试
	for i := 0; i < length; i++ {
		sli[i] = rand.Intn(1000000)
	}
	startTime = time.Now().UnixMicro()
	fmt.Println(sli)
	utils.QuickSort(0, len(sli)-1, &sli)
	fmt.Println(sli)
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("快速排序测试执行时间，毫秒：", timeDuration)

	//startTime = time.Now().UnixMicro()
	//head := &models.StudentNode{ID: -1, Next: nil}
	//err := errors.New("")
	//for i := 1; i <= length; i++ {
	//	err = head.InsertNodeByID(&models.StudentNode{
	//		ID:    rand.Intn(1000000),
	//		Name:  "无",
	//		Score: 0,
	//		Next:  nil,
	//	})
	//	if err != nil {
	//		fmt.Println("出错了，err：", err)
	//	}
	//}
	//stopTime = time.Now().UnixMicro()
	//timeDuration = stopTime - startTime
	//fmt.Println("执行时间InsertNodeByID，毫秒：", timeDuration)
	//head.ListNodes()
	//fmt.Println(sli)
}
