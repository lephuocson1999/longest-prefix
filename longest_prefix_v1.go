package main

import (
	"strings"
)

func GetLongestPrefixV1(input []string) string {
	longestPrefixMap := make(map[string]int)
	for _, v := range input {
		SetLongestPrefix(v, longestPrefixMap)
	}

	var (
		key string
	)
	for k, v := range longestPrefixMap {
		if v == len(input) && len(k) > len(key) {
			key = k
		}
	}

	if key == "" {
		return "Không có chuỗi nào"
	}

	return key
}

func SetLongestPrefix(prefix string, longestPrefixMap map[string]int) {
	// Split characters
	splitPrefix := strings.Split(prefix, "")

	var (
		longestPrefix string
	)
	for _, v := range splitPrefix {
		// If longest prefix don't contains, add to longest prefix variable
		if !strings.Contains(longestPrefix, v) {
			longestPrefix += v
			longestPrefixMap[longestPrefix]++
			continue
		}

		// Get index of v in longest prefix
		// Example: prefix = "abc" and v = "a"
		// 			lastIndex = 0
		lastIndex := strings.LastIndex(prefix, string(v))

		// Only get value from lastIndex + 1
		// Example: prefix = "abc" and v = "a"
		// 			new prefix = "bca"
		longestPrefix = longestPrefix[lastIndex+1:] + v
		longestPrefixMap[longestPrefix]++
	}
}
