package root

//User describe el objeto de un usuario
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//UserService interfaz para el servicio persistencia de un usuario
type UserService interface {
	GetUser(id int) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(User) error
	DeleteUser(id int) error
}

func main() {

}
