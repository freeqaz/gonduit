package responses

import "github.com/etcinit/gonduit/entities"

// DifferentialQueryResponse is the response of calling differential.query.
type DifferentialQueryResponse map[string]*entities.DifferentialRevision
