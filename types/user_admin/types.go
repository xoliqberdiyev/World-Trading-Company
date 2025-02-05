package types_user_admin

type UserStore interface {
	CreateUser(user *UserCreatePayload) error
	GetUserById(id string) (*UserDetailPayload, error)
	GetUserByUsername(username string) (*User, error)
	DeleteUserById(id string) error
	UpdateUserById(user *UserUpdatePayload, id string) error
	GetAllUsers() ([]*UserDetailPayload, error)
}

type User struct {
	Id        string
	Username  string
	FirstName string
	LastName  string
	Password  string
	CreatedAt string
}

type UserCreatePayload struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type UserDetailPayload struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
}

type UserUpdatePayload struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
