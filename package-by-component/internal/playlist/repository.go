package playlist

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

func (r *repositoryImpl) Create(playlist *playlist) (string, error) {
	return r.database.Create(&playlist)
}

func (r *repositoryImpl) Get(id string) (*playlist, error) {
	var playlist playlist
	return &playlist, r.database.Get("_id", id, &playlist)
}

func (r *repositoryImpl) GetAll() ([]*playlist, error) {
	var playlist []*playlist
	return playlist, r.database.GetAll(&playlist)
}
