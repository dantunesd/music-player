package user

type repository interface {
	Create(user *user) (string, error)
	GetAll() ([]*user, error)
	Get(id string) (*user, error)
}

type serviceImpl struct {
	repository repository
}

func NewService(repository repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Create(name string) (*user, error) {
	user := user{Name: name}

	id, err := s.repository.Create(&user)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

func (s *serviceImpl) Get(id string) (*user, error) {
	return s.repository.Get(id)
}

func (s *serviceImpl) GetAll() ([]*user, error) {
	return s.repository.GetAll()
}
