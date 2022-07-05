package song

type repository interface {
	Create(*song) (string, error)
	GetAll() ([]*song, error)
	Get(id string) (*song, error)
}

type serviceImpl struct {
	repository repository
}

func NewService(repository repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Create(name, artistName, albumName string, number int) (*song, error) {
	song := song{
		Name:       name,
		ArtistName: artistName,
		AlbumName:  albumName,
		Number:     number,
	}

	id, err := s.repository.Create(&song)
	if err != nil {
		return nil, err
	}

	song.ID = id
	return &song, nil
}

func (s *serviceImpl) Get(id string) (*song, error) {
	return s.repository.Get(id)
}

func (s *serviceImpl) GetAll() ([]*song, error) {
	return s.repository.GetAll()
}
