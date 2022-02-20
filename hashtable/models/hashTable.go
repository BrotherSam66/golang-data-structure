package models

import (
	"errors"
	"fmt"
)

var HashArrayLength = 7 // 哈希队列数量

// Employ 雇员的结构体
type Employ struct {
	ID   int
	Name string
	Next *Employ
}

// TODO 方法后续加

// EmpLink 指向队列的第一个节点
type EmpLink struct {
	Head *Employ
}

// TODO 方法后续加

// HashTable 定义哈希表，一个长度待定的切片
type HashTable struct {
	LinkArr []EmpLink
}

// Insert hash表插入一个数据
func (h *HashTable) Insert(employ *Employ) error {
	arrayID := HashModFun(employ.ID)    // 队列编号
	if h.LinkArr[arrayID].Head == nil { // 这个链表是空的，直接附加上
		h.LinkArr[arrayID].Head = employ
		return nil
	}
	// 到这里，Head不是空
	cur := h.LinkArr[arrayID].Head // 辅助指针，当前
	var pre *Employ = nil          // 辅助指针，cur 前面的那个
	for {
		if cur != nil {
			if cur.ID == employ.ID {
				return errors.New("已经有相同编号的节点了，插入失败！")
			} else if cur.ID > employ.ID { // 找到位置
				break
			}
			pre = cur
			cur = cur.Next

		} else { // cur==nil，说明是空Head
			break
		}
	}
	pre.Next = employ
	employ.Next = cur
	//if employ.ID < h.LinkArr[arrayID].Head.ID { // 需要插在最前面
	//	employ.Next = h.LinkArr[arrayID].Head
	//	h.LinkArr[arrayID].Head = employ
	//	return nil
	//}
	//// 到这里，不是插在Head上
	//temp := h.LinkArr[arrayID].Head
	//isNeedInsert := true // 需要插入记录
	//for {
	//	if temp.Next == nil { // 表示temp是最后一个节点，保留 isNeedInsert ，跳出后插入
	//		break
	//	}
	//	if temp.Next.ID == employ.ID { // 不允许相等
	//		return errors.New("已经有相同编号的节点了，插入失败！")
	//	} else if temp.Next.ID > employ.ID { // 本节点的下一个节点的ID,查在这里正合适
	//		break
	//	}
	//	temp = temp.Next // temp 不断向下滚仓
	//}
	//if isNeedInsert {
	//	employ.Next = temp.Next // 新节点指向游标的儿子
	//	temp.Next = employ      // 新节点加到temp尾巴上
	//}

	// TODO 到这里，表示队列不空，循环把新记录插入

	return nil
}

// List hash表显示全部数据
func (h *HashTable) List() {
	for i := 0; i < HashArrayLength; i++ {
		fmt.Printf("第%d条数据链 ： ", i)
		if h.LinkArr[i].Head == nil {
			fmt.Printf("数据链为空\n")
			continue
		}
		temp := h.LinkArr[i].Head // 游动的标尺
		for {
			fmt.Printf("{ID=%d,姓名=%s};", temp.ID, temp.Name)
			if temp.Next == nil { // 到结尾了
				fmt.Printf("\n")
				break
			}
			temp = temp.Next // 下移一个节点
		}
	}
}

// HashModFun 模运算的散列方法,以便确定在哈希表里占用哪一个队列。
func HashModFun(id int) int {
	return id % HashArrayLength
}
