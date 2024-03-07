package main

import (
	"fmt"
	"strings"
)

func GetLongestPrefixV2(input []string) string {
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return input[0]
	}

	firstValue := input[0]
	result := ""

	for _, v := range firstValue {
		if result == "" {
			result += string(v)
			continue
		}

		fmt.Println("value: ", string(v))
		for j := 1; j < len(input); j++ {
			// If contains, break out of loop to result += v
			if strings.Contains(input[j], result) || result == "" {
				continue
			}

			// Get index of v in longest prefix
			lastIndex := strings.LastIndex(result, string(v))
			fmt.Println("lastIndex", lastIndex)

			// Only get value from index + 1 of lastIndex
			result = result[lastIndex+1:] + string(v)
			fmt.Println("result 2================================================================", result)
		}

		fmt.Println("result================================================================: ", result, string(v))
		result += string(v)
	}

	return result
}
