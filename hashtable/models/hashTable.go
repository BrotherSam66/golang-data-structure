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

// Delete hash表删除一个数据
func (h *HashTable) Delete(id int) error {
	arrayID := HashModFun(id)           // 队列编号
	if h.LinkArr[arrayID].Head == nil { // 这个链表是空的
		return errors.New("ID对应链表是空的，也就是没有这个ID的数据，删除失败！")
	}
	// 到这里，Head不是空
	cur := h.LinkArr[arrayID].Head // 辅助指针，当前
	next := cur.Next               // 辅助指针，cur 后面的那个
	if next == nil {               // 只有一个头
		if cur.ID == id { // 找到，就是单独的Head
			h.LinkArr[arrayID].Head = nil // 删除
			return nil
		} else {
			return errors.New("没有这个ID的数据，删除失败！")
		}
	} else { // 说明不是只有一个头
		for cur != nil { // 循环向后
			if next == nil { // next 空，说明找不到了
				return errors.New("next 空，说明找不到了，没有这个ID的数据，删除失败！")
			}
			if next.ID == id { // 非只有一个头的链表里找到
				cur.Next = next.Next // 跳过next，相当于删除了next
				return nil
			} else if next.ID > id { // next.ID已经大于寻找id了，说明后面找不到了
				return errors.New("没有这个ID的数据，删除失败！")
			} else { // 继续向后找
				cur = cur.Next
				next = cur.Next
			}
		}

	}

	return errors.New("没有这个ID的数据，忙活到最后也没有，删除失败！")
}

// DeleteOrFind hash表删除or寻找一个数据
func (h *HashTable) DeleteOrFind(isDelete bool, id int) (Employ, error) {
	var retEmploy Employ
	arrayID := HashModFun(id)           // 队列编号
	if h.LinkArr[arrayID].Head == nil { // 这个链表是空的
		return retEmploy, errors.New("ID对应链表是空的，也就是没有这个ID的数据！")
	}
	// 到这里，Head不是空
	var pre *Employ = nil          // 辅助指针，cur 当前的前一个
	cur := h.LinkArr[arrayID].Head // 辅助指针，当前
	for cur != nil {               // 循环向后
		if pre == nil && cur.ID == id { // 链表第一个元素是要找的，（包括链表只有一个元素情形）
			if isDelete {
				h.LinkArr[arrayID].Head = cur.Next // 头指向下一个，就相当于删除
				return retEmploy, nil
			}
			return *cur, nil // 返回查询结果
		}
		if cur.ID == id { // 到这里，找到的肯定不在链表的head
			if isDelete {
				pre.Next = cur.Next // 上一个 下级指向 本节点的下一级。相当于删除本节点
				return retEmploy, nil
			}
			return *cur, nil // 返回查询结果
		}
		if cur.ID > id { // 当前cur.ID已经大于寻找id了，说明后面找不到了
			return retEmploy, errors.New("当前cur.ID已经大于寻找id了，说明后面找不到了！")
		}
		pre = cur // 标尺整体下移
		cur = cur.Next
	}
	return retEmploy, errors.New("没有这个ID的数据，忙活到最后也没有！")
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
