package playlist

type Database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}

type RepositoryImpl struct {
	Database Database
}

func (r *RepositoryImpl) Create(playlist *Playlist) (string, error) {
	return r.Database.Create(&playlist)
}

func (r *RepositoryImpl) Get(id string) (*Playlist, error) {
	var playlist Playlist
	return &playlist, r.Database.Get("_id", id, &playlist)
}

func (r *RepositoryImpl) GetAll() ([]*Playlist, error) {
	var playlist []*Playlist
	return playlist, r.Database.GetAll(&playlist)
}
