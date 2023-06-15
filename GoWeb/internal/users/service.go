package users

type Service interface {
	GetAll() ([]User, error)
	PostUser(name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error)
	PutUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error)
	PatchName(id int, name string) (User, error)
	DeleteUser(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	userList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (s *service) PostUser(name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++
	user, err := s.repository.PostUser(lastID, name, lastName, email, age, height, active, creationDate)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) PutUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error) {

	user, err := s.repository.PutUser(id, name, lastName, email, age, height, active, creationDate)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) PatchName(id int, name string) (User, error) {
	user, err := s.repository.PatchName(id, name)

	return user, err

}

func (s service) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)

	return err
}
