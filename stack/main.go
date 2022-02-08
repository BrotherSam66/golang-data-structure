package main

import (
	"errors"
	"fmt"
	"stack/models"
	"stack/utils"
	"strconv"
	"time"
)

func main() {

	//testStack()

	testCalculate()

}

// 测试栈
func testStack() {
	maxTop := 10
	length := 12 // 测试的数量,大于maxTop可以测试出错提示
	err := errors.New("")
	startTime := time.Now().UnixMicro()
	stopTime := time.Now().UnixMicro()
	timeDuration := stopTime - startTime
	var popVal int
	stack := &models.Stack{
		MaxTop: maxTop,
		Top:    -1,
	}

	// list测试
	startTime = time.Now().UnixMicro()
	err = stack.List()
	if err != nil {
		fmt.Println("出错了，err：", err)
	}
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("执行时间push，毫秒：", timeDuration)

	// push测试
	startTime = time.Now().UnixMicro()
	for i := 0; i < length; i++ {
		//err = stack.Push(rand.Intn(1000000))
		err = stack.Push(i)
		if err != nil {
			fmt.Println("出错了，err：", err)
		}
	}
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("执行时间push，毫秒：", timeDuration)

	// list测试
	startTime = time.Now().UnixMicro()
	err = stack.List()
	if err != nil {
		fmt.Println("出错了，err：", err)
	}
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("执行时间push，毫秒：", timeDuration)

	// pop测试
	startTime = time.Now().UnixMicro()
	for i := 0; i < length; i++ {
		popVal, err = stack.Pop()
		if err != nil {
			fmt.Println("出错了，err：", err)
		} else {
			fmt.Printf("Pop到值 序号=%d，取值=%d \n", i, popVal)
		}
	}
	stopTime = time.Now().UnixMicro()
	timeDuration = stopTime - startTime
	fmt.Println("执行时间push，毫秒：", timeDuration)

}

// 测试计算器
func testCalculate() {
	/*
		简易计算器 3+5*6-4
		设定2个栈，numStack存数值，operStack存运算符
		index:=0
		1. 扫描表达式，发现是一个数值，直接压栈。
		2. 如果是运算符，
		2.1. 如果 operStack 空的，直接压栈
		2.2. 如果 operStack 不空，
		2.2.1. 看 operStack 栈顶运算符优先级，如果大于等于 准备压栈的运算符。
		2.2.1.1. 就从operStack 栈pop出，再从numStack 栈pop两个数值运算，运算，结果压栈到numStack。
		2.2.1.2. （如果考虑括号、乘方、开方，超过2级别运算，还要：准备压栈的运算符，再和 operStack 当前栈顶运算符优先级比较）......

		2.2.1.3. 准备压栈的运算符压栈
		2.2.2. 看 operStack 栈顶运算符优先级，如果小于，压栈运算符。
		3. 表达式扫描完毕，依次运算符栈取出数值，数值栈取出2个数值，进行计算，运算结果压栈数值栈......直到运算符栈清空。
	*/
	maxTop := 10
	//length := 12 // 测试的数量,大于maxTop可以测试出错提示
	//err := errors.New("")
	//startTime := time.Now().UnixMicro()
	//stopTime := time.Now().UnixMicro()
	//timeDuration := stopTime - startTime
	//var popVal int
	str := "5+9*4-7-8/2+7"
	//str = "5+2*4-5"

	// 数值的栈
	numStack := &models.Stack{
		MaxTop: maxTop,
		Top:    -1,
	}

	// 运算符的栈
	operStack := &models.Stack{
		MaxTop: maxTop,
		Top:    -1,
	}

	// 扫描
	index := 0 // 扫描的辅助指针
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	for {
		ch := str[index : index+1] // ch是一个只有一个字符的字符串
		temp := int([]byte(ch)[0]) // temp是对应的ASCII码值
		if utils.IsOper(temp) {    // 是+-*/ 运算符
			if operStack.Top == -1 { // 2.1. 如果 operStack 空的，直接压栈
				operStack.Push(temp)
			} else { // 2.2. 如果 operStack 不空，

				// 2.2.1. 看 operStack 栈顶运算符优先级，如果大于等于 准备压栈的运算符。
				// 2.2.1.1. 就从operStack 栈pop出，再从numStack 栈pop两个数值运算，运算，结果压栈到numStack。
				// 2.2.1.2. （如果考虑括号、乘方、开方，超过2级别运算，还要：准备压栈的运算符，再和 operStack 当前栈顶运算符优先级比较）......
				if utils.Priority(operStack.Arr[operStack.Top]) >= utils.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = utils.Calculate(num1, num2, oper)
					numStack.Push(result) // 2.2.1.1. 结果压栈到numStack。
					operStack.Push(temp)  // 2.2.1.3. 准备压栈的运算符压栈
				} else { // 2.2.2. 看 operStack 栈顶运算符优先级，如果小于，压栈运算符。
					operStack.Push(temp)
				}
			}
		} else {
			//1. 扫描表达式，发现是一个数值，直接压栈。
			val, _ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}

		// 继续扫描
		index++
		if index >= len(str) {
			break
		}
	}

	// 3. 表达式扫描完毕，依次运算符栈取出数值，数值栈取出2个数值，进行计算，运算结果压栈数值栈......直到运算符栈清空。
	for {
		if operStack.Top < 0 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = utils.Calculate(num1, num2, oper)
		numStack.Push(result) // 结果压栈到numStack。
	}

	// 最后结果
	result, _ = numStack.Pop()
	fmt.Printf("表达式 %s = %d", str, result)

}
