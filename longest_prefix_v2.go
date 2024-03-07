package main

import (
	"strings"
)

func GetLongestPrefixV2(input []string) string {
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return input[0]
	}

	// Build longest prefix map
	longestPrefixMap := BuildLongestPrefixMap(input[0], input)

	return BuildLongestPrefix(longestPrefixMap, len(input))
}

func BuildLongestPrefixMap(firstValue string, input []string) map[string]int {
	var prefix string
	longestPrefixMap := make(map[string]int)
	for _, v := range firstValue {
		prefix = BuildPrefix(prefix, string(v))

		// Assign 1 to this prefix because it is in first index
		longestPrefixMap[prefix] = 1

		for j := 1; j < len(input); j++ {
			// If len > len of different input -> return
			if len(prefix) > len(input[j]) {
				return longestPrefixMap
			}

			// If !contains, break the loop
			if !strings.Contains(input[j], prefix) {
				break
			}

			longestPrefixMap[prefix]++
		}
	}

	return longestPrefixMap
}

func BuildPrefix(prefix, v string) string {
	// Is prefix contain character
	if !strings.Contains(prefix, string(v)) {
		prefix += string(v)
	} else {
		// Get index of v in longest prefix
		// Example: prefix = "abc" and v = "a"
		// 			lastIndex = 0
		lastIndex := strings.LastIndex(prefix, string(v))

		// Only get value from lastIndex + 1
		// Example: prefix = "abc" and v = "a"
		// 			new prefix = "bca"
		prefix = prefix[lastIndex+1:] + string(v)
	}

	return prefix
}

func BuildLongestPrefix(longestPrefixMap map[string]int, lengthInput int) string {
	var (
		key string
	)
	for k, v := range longestPrefixMap {
		// If value of key == length of input -> key is valid
		if v == lengthInput && len(k) > len(key) {
			key = k
		}
	}

	if key == "" {
		return "Không có chuỗi nào"
	}

	return key
}
