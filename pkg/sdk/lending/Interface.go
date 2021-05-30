package lending

type SDK interface {
	AllowBankAccountChange(clientName string, merchantID int) (bool, error)
}
