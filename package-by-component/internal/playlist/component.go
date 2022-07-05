package playlist

type Component interface {
	Create(userId, name string, songs []string) (*playlist, error)
	Get(id string) (*playlist, error)
	GetAll() ([]*playlist, error)
}
