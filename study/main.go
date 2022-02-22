package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup   //  计数牌
var x int               //  测试互斥锁的变量
var lock sync.Mutex     // 互斥锁
var rwlock sync.RWMutex // 读写互斥锁

// 测试第11节.指针
func modify1(x *int) {
	*x = 100
}

// 测试 第24章 并发之goroutine
func testGoroutine() {
	fmt.Println("testGoroutine....")
}
func testGoroutine2() {
	fmt.Println("testGoroutine222....")
	wg.Done() // 计数牌-1
}

func chanSender(ch chan<- int) { //  <- 表示单向通道
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func chanReceiver(chanSender <-chan int, chanReceiver chan<- int) { //  <- 表示单向通道
	for {
		temp, ok := <-chanSender
		if !ok {
			break
		}
		chanReceiver <- temp * temp
	}
	close(chanReceiver)
}

func lockAdd() {
	defer wg.Done() // 延迟执行
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

func main() {
	fmt.Println("go 语言学习练习。")
	runtime.GOMAXPROCS(2) // 指定最多几个核心参与任务。例如，后台统计，指定单个核心，防止阻塞网站运行

	// 第11节.指针： &取地址，*根据地址取值
	a := 1
	modify1(&a)
	fmt.Println(a) // 100

	// new 和 make
	var b *int
	b = new(int) // 这里才给分配内存
	*b = 100
	fmt.Println(*b)
	// make 用于slice map chan （channel）
	var c map[string]int
	c = make(map[string]int, 10) // 这里初始化
	c["小王子年龄"] = 15
	fmt.Println(c)

	// 第24章 并发之goroutine 并发：同一【时间段】内执行多个任务。并行：同一【时刻】执行多个任务
	// runtime.GOMAXPROCS(2) // 指定最多几个核心参与任务。例如，后台统计，指定单个核心，防止阻塞网站运行
	// channel，goroutine之间实现通信。
	//go testGoroutine() // go并发调用，和后面的不一定谁先执行，也可能子程序执行不了
	//fmt.Println("testGoroutine后面的句子")
	//time.Sleep(time.Second) // 延迟的话，就能保证等待子程序执行

	wg.Add(1)           // 计数牌+1
	go testGoroutine2() // go并发调用，和后面的不一定谁先执行，也可能子程序执行不了
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("用一个匿名函数执行一个goroutine。同时还是一个闭包，用了函数外面的变量i", i)
			wg.Done()
		}()
		wg.Add(1)
		go func(j int) {
			fmt.Println("这个不是闭包", j)
			wg.Done()
		}(i)
	}
	fmt.Println("testGoroutine后面的句子")
	//time.Sleep(time.Second) // 延迟的话，就能保证等待子程序执行
	wg.Wait() // 等计数牌清零，确保所有进程执行完毕

	// 第25章 并发之channel。函数之间交换数据的需要
	// CSP communicating Sequential Processes 通过通信共享内润。不是通过共享内存实现通信
	//var ch1 chan int   // 传递int的通道。引用类型，需要make初始化后才能用
	//var ch2 chan bool  // 传递bool的通道
	//var ch3 chan []int // 传递int切片的通道
	//ch4 := make(chan int)     // make初始化,无缓冲区，同步通道，必须手把手交接
	ch5 := make(chan int, 15) // 缓冲=15
	// 发送|接收 <- ；关闭
	ch5 <- 10 // 把10发送到ch5
	ch5 <- 18 // 发送到ch5
	fmt.Printf("通道中元素数量 =%d\n", len(ch5))
	fmt.Printf("通道总容量 =%d\n", cap(ch5))
	d := <-ch5 // 从ch5 接收一个数值
	fmt.Printf("d的值:=%d\n", d)
	<-ch5      // 从ch5 接收一个数值直接废弃
	close(ch5) // 关闭通道
	// ①对\关闭的通道发值panic。②关闭通道依然可以取值。③对关闭且没有值的通道取值会取到0值。④关闭已经关闭的通道panic。
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go chanSender(ch1)
	go chanReceiver(ch1, ch2)
	// 另外一种for range 取值方式
	for ret := range ch2 {
		fmt.Println(ret)
	}

	// TODO worker pool （goroutine 池）

	// select 多路复用，自动到多个通道去取值,类似于switch
	ch3 := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case data := <-ch3:
			fmt.Println("取到data == ", data)
		case ch3 <- i:
		default:
			fmt.Println("啥都不干！！")
		}
	}

	// 第26章，并发安全和锁
	wg.Add(2)
	go lockAdd()
	go lockAdd()
	wg.Wait()
	fmt.Println("x加锁后得数是稳定的：", x)
	// 读写互斥锁
	rwlock.RLock()
	rwlock.RUnlock()
	rwlock.Lock()
	rwlock.Unlock()
	// TODO 	var loadOnce sync.Once
	// loadOnce.DO(子程序) // 保证子程序只执行一次

	// TODO sync.Map 因为原生的map在并发里是不安全的，
	// 原生map可以加互斥锁，也可以用sync.map,Store存/Load取/LoadOrStore先获取如果不存在就存/Delete删除/Range遍历
	var m = sync.Map{} // keu value都是空接口，可以是任何类型
	m.Store("age", 18)
	v, _ := m.Load("age")
	fmt.Println("取得的值：=", v)

	// 第19章 接口
	d1 := dog{}
	do(d1)

}

// 第19章 接口 是一种类型，抽象的类型。其实是协议/规则 接口不管你是什么类型，只管你实现什么方法
type sayer interface {
	say()
}

func do(arg sayer) {
	arg.say()
}

type dog struct {
}

func (d dog) say() {
	fmt.Println("狗叫了，汪汪汪")

}
