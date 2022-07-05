package user

type Component interface {
	Create(name string) (*user, error)
	Get(id string) (*user, error)
	GetAll() ([]*user, error)
}
