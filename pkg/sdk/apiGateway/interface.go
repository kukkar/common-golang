package apiGateway

type SDK interface {
	StartTxn(req StartTxnReq, clientName string) (*StartTxnRes, error)
	CheckTxn(req CheckTxnReq, clientName string) (*CheckTxnRes, error)
}
