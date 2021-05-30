package payout

type SDK interface {
	ClubMember(merchantID int, clientName string) (bool, error)
}
