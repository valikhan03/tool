package tool

import "github.com/google/uuid"

// CREATE TABLE IF NOT EXISTS products(
//     id SERIAL,
//     title TEXT,
//     description TEXT,
//     start_price integer,
//     params JSONB,
//     auction_id UUID REFERENCES auctions(auction_id)
// );

type Lot struct {
	ID          uuid.UUID              `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	StartPrice  int                    `json:"start_price"`
	Params      map[string]interface{} `json:"params"`
}
