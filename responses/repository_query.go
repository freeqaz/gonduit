package responses

import "github.com/etcinit/gonduit/entities"

// RepositoryQueryResponse is the result of repository.query operations.
type RepositoryQueryResponse map[string]*entities.Repository
