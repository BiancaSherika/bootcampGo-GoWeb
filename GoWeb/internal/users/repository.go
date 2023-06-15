package users

import "fmt"

type User struct {
	ID           int
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       float64
	Active       bool
	CreationDate string
}

var userList []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	LastID() (int, error)
	PostUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error)
	PutUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error)
	PatchName(id int, name string) (User, error)
	DeleteUser(id int) error
}

type repository struct{
	db 
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	var userList []User

	err := r.db.Read(&userList)

	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) PostUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error) {
	user := User{id, name, lastName, email, age, height, active, creationDate}
	userList = append(userList, user)
	lastID = user.ID
	return user, nil
}

func (r *repository) PutUser(id int, name, lastName, email string, age int, height float64, active bool, creationDate string) (User, error) {
	user := User{
		Name:         name,
		LastName:     lastName,
		Email:        email,
		Age:          age,
		Height:       height,
		Active:       active,
		CreationDate: creationDate,
	}
	updated := false
	for i := range userList {
		if userList[i].ID == id {
			user.ID = id
			userList[i] = user
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d não encontrado", id)
	}

	return user, nil
}

func (repository) PatchName(id int, name string) (User, error) {
	var u User
	updated := false
	for i := range userList {
		if userList[i].ID == id {
			userList[i].Name = name
			updated = true
			u = userList[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("usuário %d não encontrado", id)
	}
	return u, nil

}

func (repository) DeleteUser(id int) error {
	deleted := false
	var index int
	for i := range userList {
		if userList[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("usuário %d não encontrado", id)
	}

	userList = append(userList[:index], userList[index+1:]...)
	return nil
}
