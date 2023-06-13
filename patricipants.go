package tool

type AttemptRequirements struct {
	ID             int    `db:"id"`
	AuctionID      string `db:"auction_id"`
	Approve        int    `db:"approve_required"`
	EnterFee       int    `db:"enter_fee_required"`
	EnterFeeAmount int    `db:"enter_fee_amount"`
}
