package loyalty

type SDK interface {
	GetLoyaltyPoints(merchantID int, clientName string) (*GetLoyaltyPointsRes, error)
	RedeemLoyaltyPoints(req RedeemPointsReq, clientName string) error
}
