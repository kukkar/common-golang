package kyc

type SDK interface {
	GetGSTDetail(merchantID int, clientName string) (*GSTDetail, error)
	GetPanDetail(merchantID int, clientName string) (*PANDetail, error)
	PanFetchNVerify(panCardNumber string, panName string,
		clientName string) (*PanDataVerfiy, error)
	SubmitKycDoc(req SubmitKycDocReq, clientName string) error
}
