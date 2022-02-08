package models

import (
	"errors"
	"fmt"
) 

type StudentNode struct {
	ID    int
	Name  string
	Score int
	Pre   *StudentNode
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
	newNode.Pre = temp  // 新节点上线指向temp

	return nil
}

// InsertNodeByID 按ID加入一个节点
func (n *StudentNode) InsertNodeByID(newNode *StudentNode) error {

	//if n.Next == nil { // 表示是空头
	//	n.Next = newNode
	//	newNode.Pre = n
	//	return nil
	//}

	flag := true   // 是否应该插入的标识
	isEnd := false // 是否是最后一个
	temp := n      // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			isEnd = true
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
		// 下面4句话，执行顺序很重要
		newNode.Next = temp.Next // 新节点指向游标的儿子
		if !isEnd {
			temp.Next.Pre = newNode // 加前向指针，如果本来是尾巴上加，就不用
		}
		temp.Next = newNode // 新节点加到temp尾巴上
		newNode.Pre = temp  // 加前向指针
	}
	return nil
}

// DeleteByID 按ID删除一个节点
func (n *StudentNode) DeleteByID(id int) error {
	//flag := true // 是否应该插入的标识
	//isEnd := false // 是否是最后一个
	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			return errors.New("没有找到这个ID！")
		}
		if temp.Next.ID == id { // 找到了，下一个就是
			if temp.Next.Next != nil {
				temp.Next.Next.Pre = temp // 孙子 向上 指向 我(和下一句执行先后顺序很重要)
			}
			temp.Next = temp.Next.Next // 当前节点，跳指向孙子节点。相当于删除了儿子节点(和上一句执行先后顺序很重要)
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

// ListNodesBack 反向列表全部节点
func (n *StudentNode) ListNodesBack() error {
	if n.Next == nil {
		fmt.Printf("链表是空的！！ \n")
	}
	// 先 游标到尾巴
	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			break
		}
		temp = temp.Next // temp 不断向下滚仓
	}
	// 反向打印
	for {
		if temp.Pre == nil { // 表示是head
			break
		}
		fmt.Printf("[%v, %v, %v] --> ", temp.ID, temp.Name, temp.Score)
		temp = temp.Pre // temp 不断向上滚仓
	}
	fmt.Println()
	return nil
}
