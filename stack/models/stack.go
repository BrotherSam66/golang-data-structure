package models

import (
	"errors"
	"fmt"
)

// Stack 栈
type Stack struct {
	MaxTop int
	Top    int
	Arr    [100000]int
}

// Push 入栈
func (s *Stack) Push(val int) error {
	if s.Top >= s.MaxTop-1 { // 表示溢出
		return errors.New("溢出了！")
	}
	// 不溢出
	s.Top++
	s.Arr[s.Top] = val
	return nil
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if s.Top < 0 { // 表示到栈底了
		return 0, errors.New("到栈底了！")
	}
	// 不到栈底
	s.Top--
	return s.Arr[s.Top+1], nil
}

// List 遍历
func (s *Stack) List() error {
	if s.Top < 0 { // 表示空栈
		return errors.New("空栈！")
	}
	// 不空栈
	fmt.Println("本栈的情况如下：")
	for curTop := s.Top; curTop >= 0; curTop-- {
		fmt.Printf("Arr[%d] = %d\n", curTop, s.Arr[curTop])
	}
	return nil

}
