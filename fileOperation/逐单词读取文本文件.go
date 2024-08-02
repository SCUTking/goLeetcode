package fileOperation

import (
	"bufio"
	"fmt"
	"os"
)

func lineByWord(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f) //初始化对应的扫描器
	scanner.Split(bufio.ScanWords) //设定对应的分割方式（以行做分割，以单词做分割，以字符做分割）
	var words []string
	for scanner.Scan() { //开始扫描，跟普通的文件读取类似
		words = append(words, scanner.Text()) //scanner.Text()每次分割得出的字符串
	}
	for _, word := range words {
		fmt.Println(word)
	}
	return nil
}
