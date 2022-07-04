package song

type Database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}

type RepositoryImpl struct {
	Database Database
}

func (r RepositoryImpl) Create(song *Song) (string, error) {
	return r.Database.Create(&song)
}

func (r RepositoryImpl) Get(id string) (*Song, error) {
	var song Song
	return &song, r.Database.Get("_id", id, &song)
}

func (r RepositoryImpl) GetAll() ([]*Song, error) {
	var song []*Song
	return song, r.Database.GetAll(&song)
}
