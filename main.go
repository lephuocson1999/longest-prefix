package main

import (
	"fmt"
)

// Đề bài:
// Cho 1 mảng A có n chuỗi ký tự,
// hãy tìm 1 chuỗi ký tự dài nhất mà
// chuỗi này bắt đầu ở tất cả các chuỗi đã cho trong mảng A.

// Ví dụ:
// A = [ "abc", "abf", "abcdef" ] → Output = "ab"
// A = [ "abc", "cde", "gh" ] → Output = Không có chuỗi nào
// A := []string{"aaaccccc", "aaaaaddd", "aaaagggg"} → Output = "aa"

func main() {
	A := []string{"aaaccccc", "aaaaaddd", "aaaagggg"}
	// A := []string{"abc", "abf", "abcdef"}
	// A := []string{"abc", "cde", "gh"}
	fmt.Println(GetLongestPrefixV1(A))
}
