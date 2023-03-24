package tool


type Lot struct {
	ID          int                    `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	StartPrice  int                    `json:"start_price"`
	Params      map[string]interface{} `json:"params"`
}
