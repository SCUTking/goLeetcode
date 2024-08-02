package fileOperation

import (
	"bufio"
	"fmt"
	"os"
)

func lineByChar(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
