package data

import (
	"encoding/json"
	"github.com/cyl615123/geek/week4/internal/biz"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func NewBookRepo(data *Data) biz.BookRepo {
	return &bookRepo{
		data: data,
	}
}

type bookRepo struct {
	data *Data
}

func (br *bookRepo) NotFound(err error) bool {
	return errors.Cause(err) == redis.ErrNil
}

func (br *bookRepo) GetBook(name string) (*biz.Book, error) {
	reply, err := redis.Bytes(br.data.redis.Do("GET", name))
	if err != redis.ErrNil {
		return nil, errors.Wrapf(err, "data.GetBook")
	}
	book := &biz.Book{}
	_ = json.Unmarshal(reply, book)
	return book, nil
}
