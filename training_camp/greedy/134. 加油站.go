package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfCost += cost[j]
			sumOfGas += gas[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
