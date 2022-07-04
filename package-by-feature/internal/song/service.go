package song

type Repository interface {
	Create(*Song) (string, error)
	GetAll() ([]*Song, error)
	Get(id string) (*Song, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(name, artistName, albumName string, number int) (*Song, error) {
	song := Song{
		Name:       name,
		ArtistName: artistName,
		AlbumName:  albumName,
		Number:     number,
	}

	id, err := s.Repository.Create(&song)
	if err != nil {
		return nil, err
	}

	song.ID = id
	return &song, nil
}

func (s *ServiceImpl) Get(id string) (*Song, error) {
	return s.Repository.Get(id)
}

func (s *ServiceImpl) GetAll() ([]*Song, error) {
	return s.Repository.GetAll()
}
