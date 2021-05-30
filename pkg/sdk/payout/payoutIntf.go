package payout

type SDK interface {
	PennyDrop(req PennyDropReq, clientName string) (*PennyDropRes, error)
	EarlySettlement(req ReqEarlySettlement,
		clientName string) ([]int, error)
}
