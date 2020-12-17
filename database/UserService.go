package database

import (
	"database/sql"
	"rest-api/root"
)

//UserService representa una implementación de mysql
type UserService struct {
	DB *sql.DB
}

// GetUsers busca la lista de los usuarios
func (s *UserService) GetUsers() ([]*root.User, error) {
	users := []*root.User{}
	rows, err := s.DB.Query("SELECT name FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user *root.User
		err = rows.Scan(&user.Name)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

//GetUser obtiene un user dado un id
func (s *UserService) GetUser(id int) (*root.User, error) {
	return nil, nil
}

//CreateUser crea un usuario
func (s *UserService) CreateUser() error {
	return nil
}

//DeleteUser elimina un usuario dado un id
func (s *UserService) DeleteUser(id int) error {
	return nil
}
