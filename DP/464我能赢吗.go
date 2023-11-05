package DP

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	//等差数列求和
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	m := make(map[int]bool)
	var dfs func(userNum int, currTot int) bool
	dfs = func(userNum int, currTot int) bool {
		i, ok := m[userNum]
		if !ok {
			//没有试过这个场景
			var res bool
			for i := 0; i < maxChoosableInteger; i++ {
				if ((userNum >> i) & 1) == 0 {
					if (i + 1 + currTot) >= desiredTotal {
						res = true
					}
					if !dfs((userNum | (1 << i)), currTot+i+1) {
						res = false
					}
				}
			}
			m[userNum] = res
			i = res
		}
		return i
	}
	return dfs(0, 0)
}
