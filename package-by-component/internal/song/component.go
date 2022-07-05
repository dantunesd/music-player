package song

type Component interface {
	Create(name, artistName, albumName string, number int) (*song, error)
	Get(id string) (*song, error)
	GetAll() ([]*song, error)
}
