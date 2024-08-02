package BigData

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 读取那个大的文件
func fileToHash(file string) {
	open, err := os.Open(file)
	if err != nil {

	}
	defer open.Close()

	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		cur := scanner.Text()
		//对当前IP进行hash
		writeIpToFile(cur)
	}
}

// 将字符串的ASCLL值相加，再取模
func stringToIntMod(input string, mod int) int {
	sum := 0
	for _, char := range input {
		sum += int(char)
	}
	return sum % mod
}

// 将IP追加写到文件中
func writeIpToFile(IP string) {

	mod := stringToIntMod(IP, hashValue)
	file, err := os.OpenFile(filePrefix+strconv.Itoa(mod), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {

	}
	defer file.Close()
	file.Write([]byte(IP + "\n"))
}

const filePrefix string = "tmp_"
const hashValue int = 1000

// 处理小文件列表 返回每个文件的top1的IP
func precess() []ipAndCount {
	files := getSmallFiles("./", filePrefix)
	res := make([]ipAndCount, 0, hashValue)
	for _, entry := range files {
		ipAndCountNode := processOneFile(entry)
		res = append(res, ipAndCountNode)
	}
	return res
}

type ipAndCount struct {
	ip    string
	count int
}

// 对map进行排序，输出topK个IP
func topKSort(res []ipAndCount, k int) []string {

	sort.Slice(res, func(i, j int) bool {
		return res[i].count > res[j].count
	})

	var ans []string
	if k > len(res) {
		k = len(res)
	}
	for i := 0; i < k; i++ {
		ans = append(ans, res[i].ip)
	}
	return ans
}

// 获取指定目录下的指定前缀的文件列表
func getSmallFiles(path, prefix string) []string {
	var file []string
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	for _, entry := range dir {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
			file = append(file, entry.Name())
		}
	}
	return file
}

// 处理每一个小文件，小文件可以全部载入内存，返回次数最多的一个IP
func processOneFile(file string) ipAndCount {
	oneFileIp2Count := make(map[string]int, 0)
	maxIp := ""
	open, err := os.Open(file)
	if err != nil {

	}
	defer open.Close()

	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cur := scanner.Text()
		//对当前IP进行hash
		oneFileIp2Count[cur]++
		if oneFileIp2Count[cur] > oneFileIp2Count[maxIp] {
			maxIp = cur
		}
	}
	return ipAndCount{ip: maxIp, count: oneFileIp2Count[maxIp]}
}
