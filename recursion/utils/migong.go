package utils

import "fmt"

// MiGong 迷宫
func MiGong() {
	/*
		建立二维数组表示地图
		0-》空白
		1-》墙
		2-》通路。最后是2表示走路路径。
		3-》走过，但是回退了，表示是死路
	*/

	row, column := 8, 15
	// 初始化
	myMap := make([][]int, row)
	for i := range myMap {
		myMap[i] = make([]int, column)
	}
	//fmt.Println(myMap)

	// 上下的墙
	for i := 0; i < column; i++ {
		myMap[0][i] = 1
		myMap[row-1][i] = 1
	}

	// 左右的墙
	for i := 0; i < row; i++ {
		myMap[i][0] = 1
		myMap[i][column-1] = 1
	}

	// 额外的墙
	myMap[3][1] = 1
	myMap[3][2] = 1

	myMap[row-3][column-2] = 1
	myMap[row-3][column-3] = 1
	myMap[row-3][column-4] = 1

	// 输出地图
	fmt.Println("原始的地图数组")
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Printf("%d,", myMap[i][j])
		}
		fmt.Println()
	}

	// 开始探索迷宫
	FindWay(&myMap, row, column, 1, 1)

	// 探索后输出地图
	fmt.Println("探索后的地图数组")
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Printf("%d,", myMap[i][j])
		}
		fmt.Println()
	}

}

// FindWay 发现出路
// myMap 是同一个地图，所以用引用
// row, column 总行数、列数
// i,j表示正在探索哪一个点
// 返回值 true=这个点可以探索；false=这个点不可以探索
func FindWay(myMap *[][]int, row, column, i, j int) bool {
	// 定义探索方向顺序。一个4级map数组。
	// 上下左右
	var direction = [4]map[string]int{{"row": -1, "column": 0}, {"row": 1, "column": 0}, {"row": 0, "column": -1}, {"row": 0, "column": 1}}
	// 右下左上
	direction = [4]map[string]int{{"row": 0, "column": 1}, {"row": 1, "column": 0}, {"row": 0, "column": -1}, {"row": -1, "column": 0}}
	// 下右上左
	//direction = [4]map[string]int{{"row": 1, "column": 0}, {"row": 0, "column": 1}, {"row": -1, "column": 0}, {"row": 0, "column": -1}}

	// 判断找到的方法，右下角=2就行
	if (*myMap)[row-2][column-2] == 2 {
		// 到达终点，回true就行。回逐级退回到原始的调用句柄的。
		return true
	}
	// 到这里，需要继续找
	if (*myMap)[i][j] == 0 { // 0=空白，表示是可探索的
		(*myMap)[i][j] = 2 // 这个标注为探索成功了
		if FindWay(myMap, row, column, i+direction[0]["row"], j+direction[0]["column"]) {
			return true
		} else if FindWay(myMap, row, column, i+direction[1]["row"], j+direction[1]["column"]) {
			return true
		} else if FindWay(myMap, row, column, i+direction[2]["row"], j+direction[2]["column"]) {
			return true
		} else if FindWay(myMap, row, column, i+direction[3]["row"], j+direction[3]["column"]) {
			return true
		} else { // 死路
			(*myMap)[i][j] = 3
			return false
		}

	} else { // 这个点不能探测，为1，是墙
		return false
	}
	return false
}
