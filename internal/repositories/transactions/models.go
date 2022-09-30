package transactions

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`
	Id            uuid.UUID `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        float64   `bun:"amount,notnull"`
	Currency      string    `bun:"currency,notnull"`
	CreatedAt     string    `bun:"createdat,notnull"`
	Status        bool      `bun:"status"`
}
