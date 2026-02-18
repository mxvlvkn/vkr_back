package crudmodel

type Item interface {
	TableName() string
	GetUpdateMap(any) (map[string]any, error)
}