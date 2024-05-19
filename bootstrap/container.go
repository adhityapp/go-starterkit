package bootstrap

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	ctx   context.Context
	db    *sqlx.DB
	trace *sdktrace
}
