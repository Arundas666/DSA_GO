package questions

func isValid(nums []byte) bool {
	var seen = make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if nums[i] != '.' {
			if seen[nums[i]] {
				return false
			}
			if !seen[nums[i]] {
				seen[nums[i]] = true
			}
		}
	}
	return true
}
func IsValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if !isValid(row) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		var column []byte
		for i := 0; i < 9; i++ {
			column = append(column, board[i][col])
		}
		if !isValid(column) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var subBox []byte
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					subBox = append(subBox, board[x][y])
				}
			}
			if !isValid(subBox) {
				return false
			}

		}
	}
	return true
}
