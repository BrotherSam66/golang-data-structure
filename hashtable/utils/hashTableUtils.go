package utils

import (
	"fmt"
	"hashtable/models"
	"strings"
)

// DemoHashTable 测试哈希表
func DemoHashTable() {
	/*

	 */

	//length := 12 // 测试的数量,大于maxTop可以测试出错提示
	//err := errors.New("")
	//startTime := time.Now().UnixMicro()
	//stopTime := time.Now().UnixMicro()
	//timeDuration := stopTime - startTime
	//var popVal int
	//str := "5+9*4-7-8/2+7"
	//str = "5+2*4-5"

	// 定义哈希表，队列数量在这里指定。
	hashTable := &models.HashTable{
		LinkArr: make([]models.EmpLink, models.HashArrayLength),
	}

	// 选择菜单
	for {
		var inputStr string
		var inputID int
		var inputName string
		fmt.Println("	i(insert)插入一条数据")
		fmt.Println("	f(find)寻找一条数据")
		fmt.Println("	d(delete)删除一条数据")
		fmt.Println("	l(list)显示全部数据")
		fmt.Println("	q(quit)退出")
		fmt.Println("请入字母，按回车键：")
		fmt.Scanln(&inputStr)
		if strings.ToUpper(inputStr) == "I" {
			fmt.Println("请输入雇员ID，按回车键：")
			fmt.Scanln(&inputID)
			fmt.Println("请输入雇员Name，按回车键：")
			fmt.Scanln(&inputName)
			emp := &models.Employ{ID: inputID, Name: inputName}
			err := hashTable.Insert(emp)
			if err != nil {
				fmt.Println("出错了：", err)
			}
			fmt.Println("插入结束")
			//} else if strings.ToUpper(inputStr) == "F" {
			//	fmt.Println("请输入雇员ID，按回车键：")
			//	fmt.Scanln(&inputID)
			//	hashTable.Find(inputID)
			//} else if strings.ToUpper(inputStr) == "D" {
			//	fmt.Println("请输入雇员ID，按回车键：")
			//	fmt.Scanln(&inputID)
			//	hashTable.Delete(inputID)
		} else if strings.ToUpper(inputStr) == "L" {
			hashTable.List()
		} else if strings.ToUpper(inputStr) == "Q" {
			fmt.Println("您选择了q(quit)退出")
			return
		}
	}

}
