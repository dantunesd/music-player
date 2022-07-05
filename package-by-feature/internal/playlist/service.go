package playlist

type repository interface {
	Create(*playlist) (string, error)
	GetAll() ([]*playlist, error)
	Get(id string) (*playlist, error)
}

type serviceImpl struct {
	repository repository
}

func NewService(repository repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Create(userId, name string, songs []string) (*playlist, error) {
	playlist := playlist{
		UserID: userId,
		Name:   name,
		Songs:  songs,
	}

	id, err := s.repository.Create(&playlist)
	if err != nil {
		return nil, err
	}

	playlist.ID = id
	return &playlist, nil
}

func (s *serviceImpl) Get(id string) (*playlist, error) {
	return s.repository.Get(id)
}

func (s *serviceImpl) GetAll() ([]*playlist, error) {
	return s.repository.GetAll()
}
