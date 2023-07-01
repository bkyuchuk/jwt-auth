package daos

import "github.com/jwo_auth/internal/database"

type DataAccessObject[T any] struct {
	*database.Database
}

func NewDataAccessObject[T any]() *DataAccessObject[T] {
	return &DataAccessObject[T]{}
}

//TODO: Create .go files for each entity that deviates from base DAO.
