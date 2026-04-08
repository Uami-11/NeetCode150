package main

func isValidSudoku(board [][]byte) bool {
	for row := range 9 {
		seen := make(map[byte]struct{})
		for col := range 9 {
			if board[row][col] == '.' {
				continue
			}

			if _, seenBefore := seen[board[row][col]]; seenBefore {
				return false
			}

			seen[board[row][col]] = struct{}{}
		}
	}

	for col := range 9 {
		seen := make(map[byte]struct{})
		for row := range 9 {
			if board[row][col] == '.' {
				continue
			}

			if _, seenBefore := seen[board[row][col]]; seenBefore {
				return false
			}

			seen[board[row][col]] = struct{}{}
		}
	}

	for square := range 9 {
		seen := make(map[byte]struct{})
		for i := range 3 {
			for j := range 3 {
				row := (square/3)*3 + i
				col := (square%3)*3 + j

				if board[row][col] == '.' {
					continue
				}

				if _, seenBefore := seen[board[row][col]]; seenBefore {
					return false
				}

				seen[board[row][col]] = struct{}{}

			}
		}
	}

	return true
}
