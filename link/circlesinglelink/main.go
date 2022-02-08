package main

import (
	"circlesinglelink/models"
	"errors"
	"fmt"
)

func main() {
	var err = errors.New("")
	head := &models.StudentNode{Next: nil}

	node1 := &models.StudentNode{
		ID:    5,
		Name:  "王五",
		Score: 88,
		Next:  nil,
	}
	node2 := &models.StudentNode{
		ID:    2,
		Name:  "刘二",
		Score: 100,
		Next:  nil,
	}
	node3 := &models.StudentNode{
		ID:    3,
		Name:  "张三",
		Score: 100,
		Next:  nil,
	}
	node4 := &models.StudentNode{
		ID:    6,
		Name:  "赵六",
		Score: 100,
		Next:  nil,
	}
	node5 := &models.StudentNode{
		ID:    4,
		Name:  "李四",
		Score: 59,
		Next:  nil,
	}
	//head.InsertNodeRear(node1)
	//head.InsertNodeRear(node2)
	//head.InsertNodeRear(node3)
	//head.InsertNodeRear(node4)
	//head.InsertNodeRear(node5)
	//head.ListNodes()
	//
	//head, err = head.DeleteByID(5)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//head, err = head.DeleteByID(4)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//
	//head, err = head.DeleteByID(3)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//
	//head, err = head.DeleteByID(2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//
	//head, err = head.DeleteByID(6)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//
	//head, err = head.DeleteByID(3)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//head.ListNodes()
	//
	//return

	head, err = head.InsertNodeByID(node1)
	if err != nil {
		fmt.Println(err)
	}
	head, err = head.InsertNodeByID(node2)
	if err != nil {
		fmt.Println(err)
	}
	head, err = head.InsertNodeByID(node3)
	if err != nil {
		fmt.Println(err)
	}
	head, err = head.InsertNodeByID(node4)
	if err != nil {
		fmt.Println(err)
	}
	head, err = head.InsertNodeByID(node5)
	if err != nil {
		fmt.Println(err)
	}

	head.ListNodes()

	head, err = head.DeleteByID(4)
	if err != nil {
		fmt.Println(err)
	}
	head.ListNodes()

	head, err = head.DeleteByID(4)
	if err != nil {
		fmt.Println(err)
	}
	head.ListNodes()

	// 以下为 josephu 约瑟夫问题
	head = &models.StudentNode{Next: nil}
	for i := 1; i <= 15; i++ {
		head, err = head.InsertNodeByID(&models.StudentNode{
			ID:    i,
			Name:  "无",
			Score: 0,
			Next:  nil,
		})
	}
	head.ListNodes()

	err = head.DeleteByJosephu(3, 6)
	if err != nil {
		fmt.Println(err)
	}
}
