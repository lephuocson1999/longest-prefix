package main

func GetLongestPrefixV1(input []string) string {
	if len(input) == 0 {
		return ""
	}
	// return if len của input == 1
	if len(input) == 1 {
		return input[0]
	}

	firstInput := input[0]

	// For các ký tự của chuỗi đầu tiên
	for i, character := range firstInput {
		for j := 1; j < len(input); j++ {
			// Nếu i > length của các chuỗi tiếp theo -> chuỗi đầu tiên dài hơn chuỗi còn lại -> return ra số ký tự trùng
			// Nếu kí tự ở index i của các chuỗi tiếp theo != character -> chuỗi prefix trùng nhau kết thúc -> return ra số ký tự trùng
			if i >= len(input[j]) || string(input[j][i]) != string(character) {
				if input[0][:i] == "" {
					return "Không có chuỗi nào"
				}

				return input[0][:i]
			}
		}
	}

	return "Không có chuỗi nào"
}
