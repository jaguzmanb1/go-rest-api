package data

import (
	"database/sql"
	"encoding/json"
	"io"
)

// User define la estructura de un usuario para el API
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FromJSON devuelve el objeto User como un JSON
func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

//Users is una colección de User
type Users []*User

// ToJSON Serializa el contenido de la coleccion a JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

//UserService representa una implementación de mysql
type UserService struct {
	DB *sql.DB
}

// GetUsers retorna una lista de usuarios
func (s *UserService) GetUsers() (Users, error) {
	users := Users{}
	rows, err := s.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

//GetUser obtiene un user dado un id
func (s *UserService) GetUser(id int) (User, error) {
	user := User{}
	rows, err := s.DB.Query("SELECT id, name FROM users WHERE id = (?)", id)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		user = User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return user, err
		}

		return user, err
	}

	return user, err
}

//CreateUser crea un usuario
func (s *UserService) CreateUser(pUser User) error {
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
