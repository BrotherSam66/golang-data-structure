package models

import (
	"errors"
	"fmt"
)

type StudentNode struct {
	ID    int
	Name  string
	Score int
	Next  *StudentNode
}

// InsertNodeRear 队尾加入一个节点
func (n *StudentNode) InsertNodeRear(newNode *StudentNode) error {

	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			break
		}
		temp = temp.Next // temp 不断向下滚仓
	}
	temp.Next = newNode // 新节点加到尾巴上

	return nil
}

// InsertNodeByID 按ID加入一个节点
func (n *StudentNode) InsertNodeByID(newNode *StudentNode) error {

	if n.Next == nil { // 表示是空头
		n.Next = newNode
		return nil
	}

	flag := true // 是否应该插入的标识
	temp := n    // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			break
		}
		if temp.Next.ID > newNode.ID { // 表示本节点的下一节点ID大于拟增加节点，插在这里正合适
			break
		} else if temp.Next.ID == newNode.ID { // 不允许相等
			return errors.New("已经有相同编号的节点了，插入失败！")
		}
		temp = temp.Next // temp 不断向下滚仓
	}
	if flag {
		newNode.Next = temp.Next // 新节点指向游标的儿子
		temp.Next = newNode      // 新节点加到temp尾巴上
	}

	return nil
}

// DeleteByID 按ID删除一个节点
func (n *StudentNode) DeleteByID(id int) error {
	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			return errors.New("没有找到这个ID！")
		}
		if temp.Next.ID == id { // 找到了，下一个就是
			temp.Next = temp.Next.Next // 当前节点，跳指向孙子节点。相当于删除了儿子节点
			break
		}
		temp = temp.Next // temp 不断向下滚仓
	}
	return nil
}

// ListNodes 列表全部节点
func (n *StudentNode) ListNodes() error {
	if n.Next == nil {
		fmt.Printf("链表是空的！！ \n")
	}
	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			break
		}
		fmt.Printf("[%v, %v, %v] --> ", temp.Next.ID, temp.Next.Name, temp.Next.Score)
		temp = temp.Next // temp 不断向下滚仓
	}
	fmt.Println()
	return nil
}
