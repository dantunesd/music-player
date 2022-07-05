package user

type database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}

type repositoryImpl struct {
	database database
}

func NewRepository(database database) *repositoryImpl {
	return &repositoryImpl{
		database: database,
	}
}

func (r *repositoryImpl) Create(user *user) (string, error) {
	return r.database.Create(&user)
}

func (r *repositoryImpl) Get(id string) (*user, error) {
	var user user
	return &user, r.database.Get("_id", id, &user)
}

func (r *repositoryImpl) GetAll() ([]*user, error) {
	var user []*user
	return user, r.database.GetAll(&user)
}
