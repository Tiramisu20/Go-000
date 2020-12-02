package main

import (
	// "errors"
	"fmt"
	"github.com/pkg/errors"
)
// 题目：
// 我们在操作数据库时，比如 DAO 层中遇到一个 sql.ErrNoRows 的时候，是否应该 Warp 这个 error
// 抛给上层。为什么，应该怎么做？请写出代码

func main() {
	id := 10086
	result, err := doApi(id)

	// main 函数处理全局 error
	if (err != nil) {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Println(result)
}

func doApi(id int) (int, error) {

	// do api logic

	return doService(id)
}

func doService (id int) (int, error){

	// do service logic

	result, err := doDao(id)
	if err != nil {
	    return result, err
	}

	// do service logic

	return result, err
}


///////////////////////////////////////////////////////////////////

func doDao(pk int) (int, error) {
	// handle logic
	// return 0, ErrNoRows

	err := ErrNoRows
	// 处理错误，获取数量为 0， 按业务逻辑为 0 处理
	if errors.Is(err, ErrNoRows) {
		return 0, nil
	}

	return 0, errors.New("some other error")
}

// SQL error
var ErrNoRows = errors.New("SQL: ERROR NO ROWS! ")


// 处理想法与思路：
// 对于没有从 DAO 层获取记录的情况，应该在 DAO 层处理掉这个 error, 并向业务层返回合适的值。
// 业务层，对于捕获的 error 直接向最上层返回。在 main 函数中统一处理。
// 使用 pkg/errors 包，打印下 error 的堆栈。