package user

type Database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}

type RepositoryImpl struct {
	Database Database
}

func (r *RepositoryImpl) Create(user *User) (string, error) {
	return r.Database.Create(&user)
}

func (r *RepositoryImpl) Get(id string) (*User, error) {
	var user User
	return &user, r.Database.Get("_id", id, &user)
}

func (r *RepositoryImpl) GetAll() ([]*User, error) {
	var user []*User
	return user, r.Database.GetAll(&user)
}
