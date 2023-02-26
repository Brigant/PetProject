package service

type AccountService struct {
	storage AccountStorage
}

func NewAccountService(storage AccountStorage) AccountService {
	return AccountService{storage: storage}
}

// TODO: implement the func.
func (a AccountService) CreateUser() error {
	return nil
}

// TODO: implement the func.
func (a AccountService) GetUser() error {
	return nil
}
