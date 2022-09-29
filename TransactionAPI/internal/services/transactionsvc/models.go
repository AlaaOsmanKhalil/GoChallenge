package transactionsvc

type CreatePld struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency" validate:"required"`
}
