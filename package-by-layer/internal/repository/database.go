package repository

type database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}
