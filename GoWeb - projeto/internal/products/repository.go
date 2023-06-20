package products

import (
	"fmt"
	"go-tests-aula2-morning/pkg/store"
)

type Product struct {
	Id    int
	Name  string
	Price float64
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, price float64) (Product, error)
	LastId() (int, error)
	Update(id int, name string, price float64) (Product, error)
	PartialUpdate(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastId() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id int, name string, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, name string, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{Name: name, Price: price}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) PartialUpdate(id int, name string, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	var p Product
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product
	r.db.Read(&ps)
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d não encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}
