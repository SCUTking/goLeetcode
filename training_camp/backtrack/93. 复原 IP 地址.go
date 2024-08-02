package backtrack

import "strconv"

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var dfs func(num int, pre string, preIndex int, index int)
	dfs = func(num int, pre string, preIndex int, index int) {
		if num == 0 {
			if index == len(s)+1 {
				s1 := string(pre)
				res = append(res, s1)
			} else {
				return
			}

		}
		if index == len(s)+1 {
			return
		}
		s2 := s[preIndex:index]

		dfs(num, pre, preIndex, index+1)
		if len(s2) > 1 && s2[0] == '0' {
			return
		}
		atoi, err := strconv.Atoi(s2)
		if err != nil {
			return
		} else {
			if atoi < 256 {
				pre += s2
				if num != 1 {
					pre += "."
				}

			} else {
				return
			}

		}
		preIndex = index

		num--
		dfs(num, pre, preIndex, index+1)

	}

	dfs(4, "", 0, 1)
	return res
}
