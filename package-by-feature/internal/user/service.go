package user

type Repository interface {
	Create(user *User) (string, error)
	GetAll() ([]*User, error)
	Get(id string) (*User, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(name string) (*User, error) {
	user := User{Name: name}

	id, err := s.Repository.Create(&user)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

func (s *ServiceImpl) Get(id string) (*User, error) {
	return s.Repository.Get(id)
}

func (s *ServiceImpl) GetAll() ([]*User, error) {
	return s.Repository.GetAll()
}
