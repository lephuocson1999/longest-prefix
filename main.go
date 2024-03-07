package main

import (
	"fmt"
	"log"
	"time"
)

// Đề bài:
// Cho 1 mảng A có n chuỗi ký tự,
// hãy tìm 1 chuỗi ký tự dài nhất mà
// chuỗi này bắt đầu ở tất cả các chuỗi đã cho trong mảng A.

// Ví dụ:
// A = [ "abc", "abf", "abcdef" ] → Output = "ab"
// A = [ "abc", "cde", "gh" ] → Output = Không có chuỗi nào

// A = [ "abcbac", "abf", "abcdef" ] → Output = "ab"

func main() {
	// A := []string{"abc", "abf", "abcdef"}
	A := []string{"abc", "cde", "gh"}

	start := time.Now()
	fmt.Println(GetLongestPrefixV2(A))
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
