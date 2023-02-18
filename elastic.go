package tool

const (
	AuctionsIDX = "auctions"
	LotsIDX     = "lots"
)

const (
	CANCELLED   = "CANCELLED"
	ACTIVE      = "ACTIVE"
	NOT_STARTED = "NOT STARTED"
)

const (
	CRE_AUC_EVENT = "CREATE_AUCTION"
	UPD_AUC_EVENT = "UPDATE_AUCTION"
	DEL_AUC_EVENT = "DELETE_AUCTION"

	ADD_PAR_EVENT = "ADD_PARTICIPANT"
	DEL_PAR_EVENT = "DELETE_PARTICIPANT"

	ADD_LOT_EVENT = "ADD_LOT"
	UPD_LOT_EVENT = "UPDATE_LOT"
	DEL_LOT_EVENT = "DELETE_LOT"
)

var EventFuncs = map[string]string{
	CRE_AUC_EVENT: "CreateAuction",
	UPD_AUC_EVENT: "UpdateAuction",
	DEL_AUC_EVENT: "DeleteAuction",
	ADD_PAR_EVENT: "AddParticipant",
	DEL_PAR_EVENT: "DeleteParticipant",
	ADD_LOT_EVENT: "AddLot",
	UPD_LOT_EVENT: "UpdateLot",
	DEL_LOT_EVENT: "DeleteLot",
}
