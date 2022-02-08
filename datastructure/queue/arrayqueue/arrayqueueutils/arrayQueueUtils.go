package arrayqueueutils

import (
	"arrayqueue/models"
	"fmt"
)

// GetOne 呼叫一个用户
func GetOne(queue *models.ArrayQueue) (err string) { // 参数是指针
	arrLen := len((*queue).Arr) // 加*表示是指针指向的对象，
	front := (*queue).Front
	rear := (*queue).Rear

	if rear == front {
		return "队列里没有排队数据。"
	}
	if front < 0 || front >= arrLen {
		return "队列头指针front超范围。"
	}
	if rear < 0 || rear >= arrLen {
		return "队列头指针rear超范围。"
	}
	front++
	if front == arrLen { // 如果 头front 溢出切片尾巴，尾rear绕回切片头。
		front = 0
	}
	(*queue).Front = front // 写回最新尾巴指针
	return
}

// AddOne 增加一个等待用户
func AddOne(queue *models.ArrayQueue) (err string) {
	arrLen := len((*queue).Arr) // 加*表示是指针指向的对象，
	front := (*queue).Front
	rear := (*queue).Rear
	nowNum := (*queue).Arr[rear] + 1 // 准备插入的编号是队尾数值+1

	if nowNum < 1 {
		return "派发的号码不能小于1。"
	}
	if front < 0 || front >= arrLen {
		return "队列头指针front超范围。"
	}
	if rear < 0 || rear >= arrLen {
		return "队列头指针rear超范围。"
	}
	if (front == 0 && rear == arrLen-1) || (rear+1 == front) { // （头在front切片头，尾rear在切片未）或（尾巴在头的前一个）
		return "队列满了，加不了人了。"
	}
	//if nowNum > 1 { // 刚开局，nowNum==1，头尾指针都是0，刚开局，尾指针不应该后移。（即 >1 才后移）
	rear++
	//}
	if rear == arrLen { // 如果 尾rear溢出切片尾巴，尾rear绕回切片头。
		rear = 0
	}
	(*queue).Rear = rear        // 写回最新尾巴指针
	(*queue).Arr[rear] = nowNum // 写回最新加入的编号
	return

}

// ListALl 列表所有等待用户
func ListALl(queue *models.ArrayQueue) (err string) {
	arrLen := len((*queue).Arr) // 加*表示是指针指向的对象，
	front := (*queue).Front
	rear := (*queue).Rear

	if rear == front {
		return "队列里没有排队数据。"
	}
	if front < 0 || front >= arrLen {
		return "队列头指针front超范围。"
	}
	if rear < 0 || rear >= arrLen {
		return "队列头指针rear超范围。"
	}
	var realLen int // 真实排队人数
	if rear > front {
		realLen = rear - front
	} else {
		realLen = rear - front + arrLen
	}
	nowPoint := front + 1 // 当前的指针的下一个才是等待的第一个
	fmt.Println()
	fmt.Println("序号\t位置\t编号")
	for i := 0; i < realLen; i++ {
		if nowPoint >= arrLen { // 当前指针越尾部边界，回卷到切片的头
			nowPoint = 0
		}
		fmt.Printf("%v\t%v\t%v", i+1, nowPoint, (*queue).Arr[nowPoint])
		fmt.Println()
		nowPoint++
	}

	(*queue).Front = front // 写回最新尾巴指针
	return

}
