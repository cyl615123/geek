package dao

import (
	"fmt"
	"github.com/pkg/errors"
)

type UserBaseInfo struct {
	ID int
	Name string
	Sex int
}

func GetUserBaseInfo(userId int) (*UserBaseInfo, error) {
	user := new(UserBaseInfo)
	querySql := fmt.Sprintf("select * from user where id=%d", userId)
	row := db.QueryRow(querySql)
	if err := row.Scan(&user.ID, &user.Name, &user.Sex); err != nil {
		return nil, errors.Wrapf(err, "QueryRow failed, sql=%s", querySql)
	}
	return user, nil
}