package main

/*
@Title
@Description  学习的demo
@Author
@Update
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup   //  计数牌
var x int               //  测试互斥锁的变量
var lock sync.Mutex     // 互斥锁
var rwLock sync.RWMutex // 读写互斥锁

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
	var b *int   // 指针定义，不会分配空间，必须new
	b = new(int) // 这里才给分配内存
	*b = 100
	fmt.Println("*b = ", *b) // 100
	// make 用于slice map chan （channel）
	var c map[string]int
	c = make(map[string]int, 1) // 这里初始化,10表示？？？？
	c["小王子年龄"] = 15
	c["大坏蛋年龄"] = 28
	fmt.Println(c)
	delete(c, "小王子年龄") // delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
	fmt.Println("删除一条map后", c)

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
	rwLock.RLock()
	rwLock.RUnlock()
	rwLock.Lock()
	rwLock.Unlock()
	// TODO 	var loadOnce sync.Once
	// loadOnce.DO(子程序) // 保证子程序只执行一次

	// TODO sync.Map 因为原生的map在并发里是不安全的，
	// 原生map可以加互斥锁，也可以用sync.map,Store存/Load取/LoadOrStore先获取如果不存在就存/Delete删除/Range遍历
	var m = sync.Map{}    // keu value都是空接口，可以是任何类型
	m.Store("age", 18)    // Store 存值
	v, _ := m.Load("age") // 取值
	fmt.Println("sync.Map{} 取得的值：=", v)
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("m.Range遍历 ：Key:%v ; Val:%v;  \n", k, v)
		return true // 终止迭代遍历时，返回 false。
	})
	m.Delete("age") // 删除值

	// 第19章 接口
	d1 := dog{}
	do(d1)

	// time.After() 定义一个，到3秒，把这个时间 从通道给出去。
	tchan := time.After(time.Second * 3)
	fmt.Println("tchan=", <-tchan) // 三秒后出 tchan= 2018-03-15 09:38:51.023106 +0800 CST m=+3.015805601

	// sleep
	time.Sleep(time.Second * 2)

	// 字符串拼接效率
	_ = "hello" + "," + "world"                          // ①直接运算符，慢，还会额外GC
	_ = fmt.Sprintf("%s,%s", "hello", "world")           // ② 较慢，无GC 格式化字符串并赋值给新串：
	str := strings.Join([]string{"hello", "world"}, "|") // ③ 快一些。hello|world
	_ = strings.Split(str, "|")                          // 分割字符穿为切片
	var buffer bytes.Buffer                              // ④ 快
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("world")
	_ = buffer.String()

	// 结构体中标记 Title string `gorm:"column:title; type: varchar(128)" json:"title" form:"title"`

	// rune 等同于int32,常用来处理unicode或utf-8字符，byte 等同于int8，常用来处理ascii字符
	var chinese = "我I我"
	fmt.Println("rune chinese length", len(chinese), []byte(chinese)) // 7 ，每个汉字3个字节？
	runeStr := []rune(chinese)
	fmt.Println("rune chinese word length", len([]rune(chinese)), runeStr) // 3 . rune切片 能自动识别汉字
	// 常量 显式类型定义： const b string = "abc" 	隐式类型定义： const b = "abc"
	const bConst = "abc"
	fmt.Println("常量 = ", bConst)

	// 循环 for 初始; 条件; 运算 { } |  for ;条件;{ } |  for { 死循环等待break }
	// for range 可以遍历数组、切片、字符串、map 及通道（channel）
	// for key, val := range coll { ... }
	rangeArray := [3]int{2, 4, 6} // 数组
	for key, val := range rangeArray {
		fmt.Printf("keyArray= %d ; valArray= %d \n", key, val)
	}
	rangeSlice := []int{2, 4, 6} // 切片
	for key, val := range rangeSlice {
		fmt.Printf("keySlice= %d ; valSlice= %d \n", key, val)
	}
	rangeMap := map[string]string{
		"a": "aa",
		"b": "bb",
	} // Map
	for key, val := range rangeMap {
		fmt.Printf("keyMap= %s ; valMap= %s \n", key, val)
	}

	// 第07章 切片slice 完整表达式 a[low : high : max] a := [5]int{1, 2, 3, 4, 5}  t := a[1:3:5]
	var sliceA []int // 切片定义
	arrayB := [5]int{1, 2, 3, 4, 5}
	sliceA = arrayB[1:3]  // 切片从数组切出来
	sliceC := sliceA[0:1] // 切片再切
	sliceD := make([]int, 20, 30)
	sliceE := []int{11, 12, 13, 14, 15, 16}
	sliceA = append(sliceA, 28, 29)            // append
	sliceA = append(sliceA, sliceE...)         // append form slice
	sliceA = append(sliceA[:2], sliceA[3:]...) // 切片中 删除元素
	copy(sliceD, sliceA)                       // 拷贝 a -> d,d容量不够就拷一部分
	fmt.Printf("sliceA=%v ; len(sliceA)=%d ; cap(sliceA)=%d\n", sliceA, len(sliceA), cap(sliceA))
	fmt.Printf("sliceC=%v ; len(sliceC)=%d ; cap(sliceC)=%d\n", sliceC, len(sliceC), cap(sliceC))
	fmt.Printf("sliceD=%v ; len(sliceD)=%d ; cap(sliceD)=%d\n", sliceD, len(sliceD), cap(sliceD))

	//使用空接口实现可以保存任意值的字典。空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	// TODO IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；
	// IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
	// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。

	// TODO 结构体；匿名结构体 var user struct{Name string; Age int}
	type addr struct {
		street   string
		postCode int
	}
	var p1 addr                                 // 基本实例化
	var p2 = new(addr)                          // 指针实例化
	p3 := &addr{street: "delphi", postCode: 18} // 指针实例化+赋值
	p4 := addr{street: "delphi", postCode: 18}  // 赋值
	fmt.Printf("%T,%T,%v,%v\n", p1, p2, p3, p4)
	// 结构体匿名字段，字段类型代替成员名字
	type Person struct {
		string
		int
	}
	// 嵌套结构体 嵌套匿名字段 先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找
	type Person2 struct {
		name string
		age  int
		addr
	}
	//JSON序列化：结构体-->JSON格式的字符串
	a2 := &addr{street: "wodada", postCode: 88}
	jsonStr, err := json.Marshal(a2)
	fmt.Println(a2, jsonStr, err)
	//JSON反序列化：JSON格式的字符串-->结构体
	a3 := &addr{}
	err = json.Unmarshal([]byte(jsonStr), a3)
	fmt.Println(a3, err)
}

// 第19章 接口 是一种类型，抽象的类型。其实是协议/规则 接口不管你是什么类型，只管你实现什么方法
// 有一个接口 sayer -> 有一个结构 dog -> 结构作为接收者有一个方法 say()
// -> 程序 do(arg sayer){} 指定了这个接口，内部就必须有say()方法 -> do(d1) 执行
type sayer interface {
	say()
}

func do(arg sayer) {
	arg.say()
}

type dog struct {
	name string
}

func NewDog(name string) *dog { // 构造函数,注意*号和取地址符号&。
	return &dog{ // 要求返回dog类型的指针，所以加取地址符号
		name: name,
	}
}

// 方法和接收者。
func (d dog) say() { // 方法和接收者。方法是作用于特定类型变量的函数。接收者可以是变量，也可以是指针 ,func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {函数体}
	d.name = "值接收者新名字" // 值接收者，修改不会影响原始数据
	fmt.Println("狗叫了，汪汪汪")
}
func (d *dog) sayPoint() { // 方法和接收者。方法是作用于特定类型变量的函数。接收者可以是变量，也可以是指针
	d.name = "指针接收者新名字" // 指针接收者，修改会影响原始数据
	fmt.Println("狗叫了，汪汪汪")
}

// 然后 d1:=NewDog("托尼")
// (*p1).say  语法糖可以简化为p1.say()

type myInt int // 定义自己的类型
func (m myInt) sayHi() {
	fmt.Println("自定义类型可以作为方法接收者")
}
