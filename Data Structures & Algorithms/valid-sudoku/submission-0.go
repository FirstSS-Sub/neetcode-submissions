func isValidSudoku(board [][]byte) bool {
	var squareFlag [3][3][9]bool
	for i := 0; i < 9; i++ {
		var gyoFlag [9]bool
		var retsuFlag [9]bool
		for j := 0; j < 9; j++ {
			squareI := i / 3
			squareJ := j / 3
			gyo := board[i][j] 
			retsu := board[j][i]
			if gyo != '.' {
				if gyoFlag[int(gyo)-'1'] || squareFlag[squareI][squareJ][int(gyo)-'1'] {
					return false
				}
				gyoFlag[int(gyo)-'1'] = true
				squareFlag[squareI][squareJ][int(gyo)-'1'] = true
			}
			if retsu != '.' {
				if retsuFlag[int(retsu)-'1'] {
					return false
				}
				retsuFlag[int(retsu)-'1'] = true
			}
		}
	}
	return true
}
