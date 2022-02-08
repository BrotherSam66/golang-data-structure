package models

type Sparse struct {
	Row int // 行号
	Bol int // 列号
	Val int // 值|首行是默认值
}
