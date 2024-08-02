package coding

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//	func Test1(t *testing.T) {
//		//输入：head = [1,2,3,4,5], n = 2
//		//输出：[1,2,3,5]
//		nums := []int{1, 2, 3, 4, 5}
//		head := &LinkedNode{val: nums[0]}
//		temp := head
//		for i := 1; i < len(nums); i++ {
//			head.next = &LinkedNode{val: nums[i]}
//			head = head.next
//		}
//		delete(temp, 2)
//
//		for temp != nil {
//			fmt.Println(temp.val)
//			temp = temp.next
//		}
//
// }
//
//	func Test2(t *testing.T) {
//		//前序遍历 preorder = [3,9,20,15,7]
//		//left->mid->right
//		//中序遍历 inorder = [9,3,15,20,7]
//		n := rebuild([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
//
//		res := []int{}
//		var dfs func(root *node)
//		dfs = func(root *node) {
//			if root == nil {
//				return
//			}
//
//			dfs(root.left)
//			res = append(res, root.val)
//			dfs(root.right)
//		}
//		dfs(n)
//
//		fmt.Println(res)
//
// }
func Test3(t *testing.T) {
	//3. 最长回文子串：
	//给你一个字符串s，找到 s 中最长的回文子串。
	//示例：输入：s = "babad" 输出："bab" 解释："aba" 同样是符合题意的答案。
	//输入：s = "cbbd" 输出："bb"
	res := findLongest("babad")
	fmt.Println(res)

}
func Test4(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	cat := make(chan int)
	dog := make(chan int)
	fish := make(chan int)

	go printCat(&wg, cat, fish)
	go printDog(&wg, dog, cat)
	go printFish(&wg, fish, dog)

	//利用 waitgroup 阻塞上面的三个任务
	wg.Wait()

}

func printCat(wg *sync.WaitGroup, cat chan int, fish chan int) {
	defer wg.Done()
	for true {
		fmt.Println("cat")
		cat <- 1
		<-fish
	}

}
func printDog(wg *sync.WaitGroup, dog chan int, cat chan int) {
	defer wg.Done()
	for true {
		<-cat
		fmt.Println("dog")
		dog <- 1
	}
}

func printFish(wg *sync.WaitGroup, fish chan int, dog chan int) {
	defer wg.Done()
	for true {
		<-dog
		fmt.Println("fish")
		fish <- 1
	}
}
func Test5(t *testing.T) {
	ints := make(chan int)

	ints <- 1

	go func() {
		a := <-ints
		fmt.Println(a)
	}()
}
func Test6(t *testing.T) {
	strChan := make(chan int, 0)
	numChan := make(chan int, 0)

	go func() {

		for i := 65; i < 65+26; i++ {
			<-strChan
			fmt.Printf("%v ", string(byte(i)))
			numChan <- 1
		}
	}()

	strChan <- 1

	go func() {

		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Printf("%v ", i)
			strChan <- 1
		}
	}()

	time.Sleep(time.Second * 3)
}

func Test7(t *testing.T) {
	print2()

}

func Test8(t *testing.T) {
	v := []int{1, 2, 3}

	for i, _ := range v {
		fmt.Println(v[i])
		v = append(v, v[i])
	}

	fmt.Println(v)
}
func Test9(t *testing.T) {
	v := make(map[int]int)
	v[1] = 1

	for i, _ := range v {
		v[i+1] = i + 1
	}

	fmt.Println(v)
}
func Test10(t *testing.T) {
	v := make(chan int, 10)
	v <- 1

	for i := range v {
		fmt.Println(i)
		v <- 2
	}

	fmt.Println(v)
}

func Test11(t *testing.T) {
	print1()
}
func Test12(t *testing.T) {
	c := make(chan int, 1)
	i := <-c
	fmt.Println(i)
}
