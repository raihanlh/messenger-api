package dependency

import "gitlab.com/raihanlh/messenger-api/internal/domain/user"

// Add repositories here
type Repositories struct {
	User user.Repository
}