package service

import (
	"github.com/cyl615123/geek/week4/internal/biz"
	"github.com/pkg/errors"
)

func NewLibraryService(book *biz.BookUseCase) *LibraryService {
	return &LibraryService{book: book}
}

type LibraryService struct {
	book *biz.BookUseCase
}

type SeriesReq struct {
	names []string
}

func (ls *LibraryService) FindSeriesBool(names []string) ([]*biz.Book, error) {
	var seriesBook []*biz.Book
	for _, name := range names {
		book, err := ls.book.GetBook(name)
		if !ls.book.NoBook(err) {
			return nil, errors.WithMessagef(err, "service.FindSeriesBool book name=%s", name)
		}
		seriesBook = append(seriesBook, book)
	}
	return seriesBook, nil
}
