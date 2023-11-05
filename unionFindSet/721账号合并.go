package unionFindSet

import (
	"sort"
)

func AccountsMerge(accounts [][]string) [][]string {

	m := make(map[string]int)
	for i1, account := range accounts {
		for i := 1; i < len(account); i++ {
			i2, ok := m[account[i]]
			//如果存在
			if ok {
				for j := 1; j < len(account); j++ {
					delete(m, account[j])
				}
				for j := 1; j < len(account); j++ {
					m[account[j]] = i2
				}
			} else {
				m[account[i]] = i1
			}
		}
	}

	res := make([][]string, len(accounts))
	for k, v := range m {
		if len(res[v]) == 0 {
			res[v] = append(res[v], accounts[v][0])
		}
		res[v] = append(res[v], k)
	}
	res1 := make([][]string, 0)
	for _, re := range res {
		if re == nil {
			continue
		} else {

			sort.Strings(re)
			res1 = append(res1, re)
		}
	}

	//i := res[3:]
	//fmt.Println(i)
	return res1

}
