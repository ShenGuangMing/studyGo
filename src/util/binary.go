package util

import "strings"

// BinaryFormat 输出一个整形数的二进制，第一位是他的符号位
func BinaryFormat(n int32) string {
	sb := strings.Builder{}
	for n != 0 {
		if n%2 == 0 {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
		n /= 2
	}
	if n < 0 {
		n *= -1
		sb.WriteString("1")
	} else {
		sb.WriteString("0")
	}
	return ""
}

// Reverse 将一个字符倒序
func Reverse(str string) string {
	return ""
}
