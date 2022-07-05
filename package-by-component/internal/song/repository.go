package song

type database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}

type repositoryImpl struct {
	database database
}

func NewRepositoryImpl(database database) *repositoryImpl {
	return &repositoryImpl{
		database: database,
	}
}

func (r *repositoryImpl) Create(song *song) (string, error) {
	return r.database.Create(&song)
}

func (r *repositoryImpl) Get(id string) (*song, error) {
	var song song
	return &song, r.database.Get("_id", id, &song)
}

func (r *repositoryImpl) GetAll() ([]*song, error) {
	var song []*song
	return song, r.database.GetAll(&song)
}
