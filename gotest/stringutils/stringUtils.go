package stringutils

// Reverse 字符串首尾反转子程序
func Reverse(str string) string {
	//n := len(str) // 字符串长度
	//fmt.Println("字符串长度为：", n)
	// 考虑到有些汉字占用多个字节，先转为rune切片。
	runes := []rune(str)
	lenRunes := len(runes)
	//fmt.Println("字符串rune长度为：", lenRunes)
	retRunes := make([]rune, lenRunes, lenRunes) // 造一个等长的切片，准备接收倒叙
	for i := lenRunes; i > 0; i-- {
		retRunes[lenRunes-i] = runes[i-1]
	}
	retStr := string(retRunes)
	//fmt.Println("返回字符串长度为：", len(retStr))
	//fmt.Println("返回字符串为：", retStr)
	return retStr
}
