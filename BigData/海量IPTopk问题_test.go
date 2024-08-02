package BigData

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fileToHash("./test.txt")
	sort := topKSort(precess(), 2)
	fmt.Println(sort)
}
