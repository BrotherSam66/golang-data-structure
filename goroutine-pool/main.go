package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 协程池

const POOLMAX = 2 // 最大协程数量
var wg sync.WaitGroup

// 方案一 读写锁
var lock sync.Mutex
var goroutineNow int // 控制正在运行的协程小于这个
// 工作01
func worker01(jobId int) {
	defer wg.Done() // 退出前执行，协程数--
	duration := rand.Intn(1000)
	fmt.Printf("jobId = %d 开始，延迟 %d毫秒\n", jobId, duration)
	time.Sleep(time.Millisecond * time.Duration(duration)) // 随机延迟1~1000毫秒
	fmt.Printf("jobId = %d 结束 \n", jobId)

	lock.Lock()
	goroutineNow-- // pool指标空闲出一个
	fmt.Println("--后的goroutineNow == ", goroutineNow)
	lock.Unlock()
}

// 方案二 Chanel
func worker02(jobId int, ch chan struct{}) {
	defer wg.Done() // 退出前执行，协程数--
	duration := rand.Intn(1000)
	fmt.Printf("jobId = %d 开始，延迟 %d毫秒\n", jobId, duration)
	time.Sleep(time.Millisecond * time.Duration(duration)) // 随机延迟1~1000毫秒
	fmt.Printf("jobId = %d 结束 \n", jobId)
	_ = <-ch // 运行结束后，通道取出一个数，让出位置，让其他协程进来

}

func main() {
	// runtime.GOMAXPROCS(8)

	// 方案一 读写锁
	goroutineNow = 0
	for i := 0; i < 0; i++ {
		for { // 死循环
			if goroutineNow < POOLMAX { // pool有空闲指标
				lock.Lock()    // 加锁
				goroutineNow++ // 改队列数值
				fmt.Println("++后的goroutineNow == ", goroutineNow)
				lock.Unlock()
				wg.Add(1)
				go worker01(i) // 协程运行业务
				break
			} else { // pool 没有空闲指标
				// time.Sleep(time.Millisecond * 100)
			}
		}
	}
	fmt.Printf("最后的 wg == %#v \n", wg)
	wg.Wait() // 等待所有协程运算完成
	fmt.Printf("最后的 wg == %#v \n", wg)
	fmt.Println("最后的goroutineNow == ", goroutineNow)

	// 方案二 Chanel
	ch := make(chan struct{}, POOLMAX) // 控制用通道,通道缓存数量，就是pool容量
	iMax := 20
	i := 0
	for i < iMax { // 循环
		select {
		case ch <- struct{}{}: // 通道还有空闲，
			wg.Add(1)
			go worker02(i, ch) // 协程运行业务
			i++
			break
		default:
			break
		}
	}
	fmt.Printf("最后的 wg == %#v \n", wg)
	wg.Wait() // 等待所有协程运算完成
	fmt.Printf("最后的 wg == %#v \n", wg)
}
