package searching

func LongestCommonSubsequence(x []byte, y []byte) (int,[]byte) {
	lenX := len(x)
	lenY := len(y)
	if lenX != lenY {
		//BugBug : this need to be fixed to support the case.
		panic("Different lengths of sequence not supported due to bug.")
	}
	c := make([][]int, lenX+1)
	for i:= range c{
		c[i] = make([]int, lenY+1)
	}
	for i := 0; i <= lenX; i++{
		for j := 0; j <= lenY; j++{
			if i == 0 || j == 0 {
				c[i][j] = 0
			} else if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
		}
	}
	return c[lenX][lenY], buildLCSSequence(c, x, y, lenX-1, lenY-1)
}

func buildLCSSequence(table [][]int, x, y []byte, i, j int) []byte {
	if i == -1 || j == -1 {
		return []byte{}
	} else if x[i] == y[j] {
		return append(buildLCSSequence(table, x, y, i-1, j-1), x[i])
	} else {
		if table[i][j-1] > table[i-1][j] {
			return buildLCSSequence(table, x, y, i, j-1)
		} else {
			return buildLCSSequence(table, x, y, i-1, j)
		}
	}
}

func LCSLength(x, y []byte) int{
	if len(x) ==0 || len(y) == 0{
		return 0
	}
	if x[len(x)-1] == y[len(y)-1]{
		return 1 + LCSLength(x[0:len(x)-1], y[0:len(y)-1])
	} else {
		return max(LCSLength(x, y[0:len(y)-1]),
			LCSLength(x[0:len(x)-1], y))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}