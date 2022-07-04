package playlist

type Repository interface {
	Create(*Playlist) (string, error)
	GetAll() ([]*Playlist, error)
	Get(id string) (*Playlist, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(userId, name string, songs []string) (*Playlist, error) {
	Playlist := Playlist{
		UserID: userId,
		Name:   name,
		Songs:  songs,
	}

	id, err := s.Repository.Create(&Playlist)
	if err != nil {
		return nil, err
	}

	Playlist.ID = id
	return &Playlist, nil
}

func (s *ServiceImpl) Get(id string) (*Playlist, error) {
	return s.Repository.Get(id)
}

func (s *ServiceImpl) GetAll() ([]*Playlist, error) {
	return s.Repository.GetAll()
}
