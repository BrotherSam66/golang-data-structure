package main

import (
	"doublelink/models"
	"fmt"
)

func main() {
	head := &models.StudentNode{Pre: nil, Next: nil}

	node1 := &models.StudentNode{
		ID:    5,
		Name:  "王五",
		Score: 88,
		Pre:   nil,
		Next:  nil,
	}
	node2 := &models.StudentNode{
		ID:    2,
		Name:  "刘二",
		Score: 100,
		Pre:   nil,
		Next:  nil,
	}
	node3 := &models.StudentNode{
		ID:    3,
		Name:  "张三",
		Score: 100,
		Pre:   nil,
		Next:  nil,
	}
	node4 := &models.StudentNode{
		ID:    6,
		Name:  "赵六",
		Score: 100,
		Pre:   nil,
		Next:  nil,
	}
	node5 := &models.StudentNode{
		ID:    4,
		Name:  "李四",
		Score: 59,
		Pre:   nil,
		Next:  nil,
	}

	//head.InsertNodeRear(node1)
	//head.InsertNodeRear(node2)
	//head.InsertNodeRear(node3)
	//head.InsertNodeRear(node4)
	//head.InsertNodeRear(node5)
	//head.ListNodes()
	//head.ListNodesBack()

	err := head.InsertNodeByID(node1)
	if err != nil {
		fmt.Println(err)
	}
	err = head.InsertNodeByID(node2)
	if err != nil {
		fmt.Println(err)
	}
	err = head.InsertNodeByID(node3)
	if err != nil {
		fmt.Println(err)
	}
	err = head.InsertNodeByID(node4)
	if err != nil {
		fmt.Println(err)
	}
	err = head.InsertNodeByID(node5)
	head.ListNodes()
	head.ListNodesBack()
	err = head.DeleteByID(4)
	if err != nil {
		fmt.Println(err)
	}
	head.ListNodes()
	head.ListNodesBack()
	err = head.DeleteByID(4)
	if err != nil {
		fmt.Println(err)
	}
	head.ListNodes()
	head.ListNodesBack()
}
