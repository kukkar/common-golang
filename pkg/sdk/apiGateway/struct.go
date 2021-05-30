package apiGateway

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Config
type Config struct {
	IPPort          string
	Version         string
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

type StartTxnReq struct {
	MerchantID int     `json:"merchantId"`
	Mid        string  `json:"mid"`
	OrderID    string  `json:"orderId"`
	Amount     float32 `json:"amount"`
}

type StartTxnRes struct {
	Status          string  `json:"status"`
	ResponseCode    string  `json:"responseCode"`
	ResponseMessage string  `json:"responseMessage"`
	Mid             string  `json:"mid"`
	TxnID           string  `json:"bharatpeTxnId"`
	CreatedAt       string  `json:"createdTimestamp"`
	Amount          float32 `json:"amount"`
	OrderID         string  `json:"orderId"`
	UPIString       string  `json:"upiString"`
}
type CheckTxnReq struct {
	Mid     string `json:"mid"`
	OrderID string `json:"orderId"`
	TxnID   string `json:"bharatpeTxnId"`
}
type CheckTxnRes struct {
	Status           string  `json:"status"`
	ResponseCode     string  `json:"responseCode"`
	ResponseMessage  string  `json:"responseMessage"`
	Mid              string  `json:"mid"`
	TxnID            string  `json:"bharatpeTxnId"`
	CreatedAt        string  `json:"createdTimestamp"`
	Amount           float32 `json:"amount"`
	OrderID          string  `json:"orderId"`
	UPIString        string  `json:"upiString"`
	PaymentStatus    string  `json:"paymentStatus"`
	BankReferenceNo  string  `json:"bankReferenceNo"`
	PaymentTxnId     *int    `json:"paymentTxnId"`
	PaymentTimestamp string  `json:"paymentTimestamp"`
}
