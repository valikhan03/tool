package tool_models

type PaymentInvoice struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	ProductType int    `json:"product_type" db:"product_type"`
	ProductID   int    `json:"product_id" db:"product_id"`
	ProductName string `json:"product_name" db:"product_name"`
	Amount      int64  `json:"amount" db:"amount"`
	Currency    string `json:"currency" db:"currency"`
	PayStatus   int    `json:"p_is_payed" db:"p_is_payed"`
}
