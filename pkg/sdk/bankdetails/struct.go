package bankdetail

import "fmt"

//Config
type Config struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	return nil
}

type SubmitBankDetail struct {
	BankCode string `json:"bank_code"`
	Mobile   string `json:"mobile"`
}
type SubmitBankDetailRes struct {
	ResponseCode string               `json:"responseCode"`
	Message      string               `json:"message"`
	Status       string               `json:"status"`
	Data         SubmitBankDetailData `json:"data"`
}

type SubmitBankDetailData struct {
	Mobile       string `json:"mobile"`
	BankDetailID string `json:"bank_details_id"`
	BankCode     string `json:"bank_code"`
}

type GetBankDetailRequest struct {
	BankDetailID string `json:"bank_details_id"`
}
type GetBankDetailRes struct {
	ResponseCode string            `json:"responseCode"`
	Message      string            `json:"message"`
	Status       string            `json:"status"`
	Data         GetBankDetailData `json:"data"`
}
type GetBankDetailData struct {
	Status          string        `json:"status"`
	BankCode        string        `json:"bank_code"`
	Mobile          string        `json:"mobile"`
	BankDetailID    string        `json:"bank_details_id"`
	AccountNumber   string        `json:"account_number"`
	IFSC            string        `json:"ifsc"`
	BeneficiaryName string        `json:"beneficiary_name"`
	AccountType     string        `json:"account_type"`
	AccountList     []AccountList `json:"account_list"`
}

type AccountList struct {
	AccountNumber   string `json:"account_number"`
	IFSC            string `json:"ifsc"`
	BeneficiaryName string `json:"beneficiary_name"`
	AccountType     string `json:"account_type"`
}
