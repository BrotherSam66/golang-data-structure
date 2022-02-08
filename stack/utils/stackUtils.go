package utils

import "fmt"

// IsOper 判断是否运算符
func IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// Calculate 运算方法本身
func Calculate(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误！")
	}
	fmt.Printf("计算 num1=%d ;num2=%d ;oper=%d ;res=%d \n", num1, num2, oper, res)
	return res
}

// Priority 计算运算符优先级
func Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}
