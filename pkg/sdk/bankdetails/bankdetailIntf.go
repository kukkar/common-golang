package bankdetail

type SDK interface {
	SubmitBankDetail(req SubmitBankDetail) (*SubmitBankDetailRes, error)
	GetBankDetail(req GetBankDetailRequest) (*GetBankDetailRes, error)
}
