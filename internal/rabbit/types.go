package rabbit

import "github.com/gitarchived/service/internal/db"

type Repository struct {
	Repository      db.Repository
	LastCommitKnown string
}
