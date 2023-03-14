package service

type DirectorService struct {
	storage DirectorStorage
}

func NewDirectorService(storage DirectorStorage) DirectorService {
	return DirectorService{storage: storage}
}

// TODO: implement the func.
func (d DirectorService) CreatDirector() error {
	return nil
}

// TODO: implement the func.
func (d DirectorService) GetDirector() error {
	return nil
}

// TODO: implement the func.
func (d DirectorService) UpdateDirector() error {
	return nil
}
