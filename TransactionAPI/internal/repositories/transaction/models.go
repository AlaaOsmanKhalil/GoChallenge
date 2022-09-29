package transaction

import (
	"github.com/uptrace/bun"
	"time"
)

type Model struct {
	bun.BaseModel `bun:"table:transactions"`

	ID        string    `bun:"id,pk"`
	Amount    float64   `bun:"amount"`
	Currency  string    `bun:"currency"`
	CreatedAt time.Time `bun:"created_at"`
}
