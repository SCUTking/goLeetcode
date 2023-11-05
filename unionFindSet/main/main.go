package main

import "goLeetcode/unionFindSet"

func main() {
	i := [][]string{}
	i = append(i, []string{"David", "David0@m.co", "David4@m.co", "David3@m.co"})
	i = append(i, []string{"David", "David5@m.co", "David0@m.co"})
	i = append(i, []string{"David", "David1@m.co", "David4@m.co", "David0@m.co"})
	i = append(i, []string{"David", "David0@m.co", "David1@m.co", "David3@m.co"})
	i = append(i, []string{"David", "David4@m.co", "David1@m.co", "David3@m.co"})
	//i = append(i, []string{"Mary", "mary@mail.com"})
	unionFindSet.AccountsMerge(i)
}
