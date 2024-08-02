package Hot100

type task struct {
	Val       uint8
	Count     int
	LastIndex int
}

func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]*task)
	for i := 0; i < len(tasks); i++ {
		cur := tasks[i]
		if _, e := m[cur]; !e {
			m[cur] = &task{
				Val:       cur,
				Count:     1,
				LastIndex: -n + 1,
			}
		} else {
			m[cur].Count++
		}
	}
	curTime := 1
	end := false
	for !end {
		maxC := -1
		var curByte byte = 0
		for k, t := range m {
			if t.Count > maxC && t.LastIndex+n >= curTime {
				maxC = t.Count
				curByte = k
			}
		}
		if curByte == 0 {
			//表示没一个符合要求的
			if len(m) == 0 {
				end = true
				break
			}
		} else {
			m[curByte].Count--
			m[curByte].LastIndex = curTime
			if m[curByte].Count == 0 {
				delete(m, curByte)
			}
		}
		curTime++
	}
	return curTime - 1
}
