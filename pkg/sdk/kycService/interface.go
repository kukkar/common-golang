package kycService

type SDK interface {
	GetPanHolderName(panCardNumber string, merchantID int, clientName string) (string, error)
	NameMatch(name1 string, name2 string, clientName string) (*int, error)
	UpdateMerchantKyc(pan string, token string, clientName string) error
	GetPan(merchantId int, clientName string) (*GetPanData, error)
}
