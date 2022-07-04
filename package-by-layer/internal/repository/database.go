package repository

type Database interface {
	Get(fieldName, fieldValue string, output interface{}) error
	Create(content interface{}) (string, error)
	GetAll(output interface{}) error
}
