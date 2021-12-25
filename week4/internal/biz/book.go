package biz

func NewBookUseCase(repo BookRepo) *BookUseCase {
	return &BookUseCase{repo: repo}
}

type Book struct {
	name string
}

type BookRepo interface {
	NotFound(err error) bool
	GetBook(name string) (*Book, error)
}

type BookUseCase struct {
	repo BookRepo
}

func (us *BookUseCase) NoBook(err error) bool {
	return us.repo.NotFound(err)
}

func (us *BookUseCase) GetBook(name string) (*Book, error) {
	return us.repo.GetBook(name)
}
