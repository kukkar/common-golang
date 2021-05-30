package swipe

type SDK interface {
	CreateAccount(req CreateAccountReq, clientName string) (*CreateAccountRes, error)
}
