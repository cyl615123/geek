package main

import (
	"database/sql"
	"fmt"
	"github.com/cyl615123/geek/week2/dao"
	"github.com/pkg/errors"
)

func main() {
	if err := dao.InitDB(); err != nil{
		fmt.Printf("initdb err=%v", err)
		return
	}
	user, err := dao.GetUserBaseInfo(123)
	if err != nil {
		fmt.Printf("GetUserBaseInfo failed, err=%+v", err)
		if sql.ErrNoRows == errors.Cause(err) {
			//业务处理
			return
		}
		//其他错误处理
		return
	}

	fmt.Printf("userbaseinfo:%#v", user)
	return
}
