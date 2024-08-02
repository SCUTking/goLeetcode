package main

import (
	"fmt"
	"sync"
)

// 数据库连接对象实例
type DBInstance struct {
}

var (
	once       sync.Once
	dbInstance *DBInstance
)

// 注意初始对象实例过程也可以放到init函数中进行
func NewInstance() *DBInstance {
	// 只有第一次才会实例化数据库连接对象
	if dbInstance == nil {
		once.Do(func() {
			dbInstance = &DBInstance{}
		})
	}
	return dbInstance
}
func main() {
	for i := 0; i <= 10; i++ {
		// 在使用的时候获取数据库实例对象
		db := NewInstance()
		fmt.Println(db)
	}

}
