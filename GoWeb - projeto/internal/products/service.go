package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, price float64) (Product, error)
	Update(id int, name string, price float64) (Product, error)
	PartialUpdate(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name string, price float64) (Product, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}

	lastId++

	product, err := s.repository.Store(lastId, name, price)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name string, price float64) (Product, error) {
	return s.repository.Update(id, name, price)
}

func (s *service) PartialUpdate(id int, name string, price float64) (Product, error) {
	return s.repository.PartialUpdate(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
