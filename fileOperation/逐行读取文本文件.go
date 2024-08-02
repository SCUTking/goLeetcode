package fileOperation

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f) //生成一个新的读取器
	for true {
		line, err := reader.ReadString('\n') //读到换行符就结束，就代表一行内容
		if err == io.EOF {                   //文件最后的字符是EOF，将以文件格式返回
			break
		} else if err != nil {
			fmt.Print(err)
		}
		fmt.Print(line) //对一行字符串进行操作
	}
	return nil
}

func lineByLine1(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
