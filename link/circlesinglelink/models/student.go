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
	if n.Next == nil { // 表示是空的
		n.ID = newNode.ID
		n.Name = newNode.Name
		n.Score = newNode.Score
		n.Next = n
		return nil
	}
	// 不为空
	newNode.Next = n.Next // 新节点嵌在头节点后面就好
	n.Next = newNode
	return nil
}

// InsertNodeByID 按ID加入一个节点
func (n *StudentNode) InsertNodeByID(newNode *StudentNode) (*StudentNode, error) {

	if n.Next == nil { // 表示是空头
		newNode.Next = newNode
		return newNode, nil
	}

	isInsert := true      // 是否应该插入的标识
	isInsertHead := false // 是否应该插入在头节点
	temp := n             // 游标 temp = 头

	if n.ID > newNode.ID { // 表示头点ID大于拟增加节点，插在头节点之前
		isInsertHead = true
	} else {
		for {
			if temp.Next == n { // 表示是最后一个节点，跳去插入
				break
			}
			if temp.Next.ID > newNode.ID { // 表示本节点的下一节点ID大于拟增加节点，插在这里正合适
				break
			} else if temp.Next.ID == newNode.ID { // 不允许相等
				return n, errors.New("已经有相同编号的节点了，插入失败！")
			}
			temp = temp.Next // temp 不断向下滚仓
		}
	}
	if isInsert { // 表示可以插入了
		if isInsertHead { // 表示插入在头节点
			for { // 一直游标跑到最后一个节点（游标的下一节点=头）
				if temp.Next == n {
					break
				}
			}
			newNode.Next = n    // 新节点指向游标的儿子
			temp.Next = newNode // 原尾巴指向新头
			return newNode, nil // 返回恰当的新头
		} else { // 表示可以插入了，新节点在游标的后面
			newNode.Next = temp.Next // 新节点指向游标的儿子
			temp.Next = newNode      // 新节点加到temp尾巴上
		}
	}

	return n, nil
}

// DeleteByID 按ID删除一个节点
func (n *StudentNode) DeleteByID(id int) (*StudentNode, error) {

	temp := n // 游标 temp = 头
	i := 0    // 用来判断是否循环起来了

	if temp.Next == nil { // 表示是空环
		return n, errors.New("这个是空环！")
	}
	for {
		if temp.Next.ID == id { // 找到了，下一个就是
			if temp.Next == temp { // 本环只有自己一个元素
				*n = StudentNode{Next: nil} // 字接清空
				return n, nil
			}
			// 找到了，且元素多于1个
			if temp.Next == n { // 找到了，下一个就是，且下一个是head
				temp.Next = temp.Next.Next // 当前节点，跳指向孙子节点。相当于删除了儿子节点
				return temp.Next, nil      // 这样保证新的头是原来的第二个
			}
			temp.Next = temp.Next.Next // 当前节点，跳指向孙子节点。相当于删除了儿子节点
			return n, nil
		}
		if temp == n && i > 0 { // 说明找回头了
			return n, errors.New("没有找到想删除的ID！")
		}
		temp = temp.Next // temp 不断向下滚仓
		i++
	}
	return n, nil
}

// DeleteByJosephu 按ID删除一个节点
func (n *StudentNode) DeleteByJosephu(fromID int, step int) error {
	fmt.Println()
	if n.Next == nil { // 表示是空环
		return errors.New("这个是空环！")
	}

	if n.Next == n { // 只有一个元素
		fmt.Printf("[%v] --> ", n.ID)
		return nil
	}

	temp := n // 游标 temp = 头
	i := 0    // 用来判断是否循环起来了

	// 先找到起始的位置
	for i = 1; i < fromID; i++ {
		temp = temp.Next
	}

	// 不断循环
	for {
		// 跳过指定的步数1
		for i = 1; i < step; i++ {
			temp = temp.Next
		}

		// 删掉temp.Next、打印、删掉
		fmt.Printf("[%v] --> ", temp.Next.ID)
		temp.Next = temp.Next.Next

		if temp.Next == temp { // 只有一个元素
			fmt.Printf("[%v] --> ", temp.ID)
			return nil
		}
	}
	return nil
}

// ListNodes 列表全部节点
func (n *StudentNode) ListNodes() error {
	fmt.Println()
	if n.Next == nil {
		fmt.Printf("链表是空的！！ \n")
	}
	temp := n // 游标 temp = 头
	for {
		if temp.Next == nil { // 表示是最后一个节点
			break
		}
		fmt.Printf("[%v, %v, %v] --> ", temp.ID, temp.Name, temp.Score)
		if temp.Next == n { // 如果下一个循环到head，停止。
			return nil
		}
		temp = temp.Next // temp 不断向下滚仓
	}
	return nil
}
