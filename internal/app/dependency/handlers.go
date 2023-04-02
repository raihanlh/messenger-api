package dependency

import (
	"gitlab.com/raihanlh/messenger-api/internal/domain/user"
	"gitlab.com/raihanlh/messenger-api/internal/health"
)

type Handlers struct {
	User   user.Handler
	Health health.Handler
}
