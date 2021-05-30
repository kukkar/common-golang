package PG

type SDK interface {
	CreateTxn(req CreateTxnReq, clientName string) (*CreateTxnData, error)
	CheckTxn(req CheckTxnReq, clientName string) (*CheckTxnResData, error)
}
