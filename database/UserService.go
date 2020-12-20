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
	rows, err := s.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := &root.User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

//GetUser obtiene un user dado un id
func (s *UserService) GetUser(id int) (*root.User, error) {
	user := &root.User{}
	rows, err := s.DB.Query("SELECT id, name FROM users WHERE id = (?)", id)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		user := &root.User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return user, err
		}

		return user, err
	}

	return nil, nil
}

//CreateUser crea un usuario
func (s *UserService) CreateUser(pUser root.User) error {
	_, err := s.DB.Exec("INSERT INTO users (name) VALUES (?)", pUser.Name)
	return err
}

//DeleteUser elimina un usuario dado un id
func (s *UserService) DeleteUser(id int) error {
	_, err := s.DB.Exec("DELETE FROM users WHERE id = (?)", id)
	if err != nil {
		return err
	}
	return nil
}
