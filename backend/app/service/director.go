package service

type DirectorService struct {
	storage DirectorStorage
}

func NewDirectorService(storage DirectorStorage) DirectorService {
	return DirectorService{storage: storage}
}

func (d DirectorService) CreatDirector() error {
	return nil
}

func (d DirectorService) GetDirector() error {
	return nil
}

func (d DirectorService) UpdateDirector() error {
	return nil
}
