package payout

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Config
type Config struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	ClientValidator clientvalidator.ClientValidator
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	if this.ClientValidator == nil {
		return fmt.Errorf("client validator required")
	}
	return nil
}

type PennyDropReq struct {
	AccountNumber string
	IFSC          string
	MerchantID    string
	MobileNumber  string
}

type PennyDropRes struct {
	IFSC            string
	BeneficiaryName string
	AccountNumber   string
}

type pennyDropServiceReq struct {
	AccountNumber string `json:"accountNo"`
	IFSC          string `json:"ifsc"`
	MerchantID    string `json:"merchantId,omitempty"`
	MobileNumber  string `json:"mobile,omitempty"`
}

type pennyDropServiceRes struct {
	ResponseCode    string          `json:"responseCode"`
	Message         string          `json:"message"`
	Status          string          `json:"status"`
	BeneficiaryData beneficiaryData `json:"data"`
}

type beneficiaryData struct {
	IFSC            string `json:"ifsc"`
	AccountNumber   string `json:"accountNo"`
	BeneficiaryName string `json:"beneficiaryName"`
}

type ReqEarlySettlement struct {
	MerchantID      int
	Amount          *float64
	MerchantStoreID *int
	Token           string
}

type serviceReqEarlySettlement struct {
	MerchantID      int      `json:"merchantId"`
	Amount          *float64 `json:"merchantStoreId,omitempty"`
	MerchantStoreID *int     `json:"amount,omitempty"`
	Token           string   `json:"token"`
}

type serviceResEarlySettlement struct {
	Status     string                     `json:"status"`
	StatusCode string                     `json:"responseCode"`
	Data       serviceDataEarlySettlement `json:"data"`
}

type serviceDataEarlySettlement struct {
	SettlementID  int   `json:"settlementId"`
	SettlementIDs []int `json:"settlementIds"`
}

type tmpservReqEarlySett struct {
	MerchantID      string  `json:"merchantId"`
	Amount          *string `json:"merchantStoreId,omitempty"`
	MerchantStoreID *string `json:"amount,omitempty"`
	Token           string  `json:"token"`
}
