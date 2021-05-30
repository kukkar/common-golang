package PG

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

type CreateTxnReq struct {
	Mid                   string   `json:"mid"`
	OrderID               string   `json:"orderId"`
	Amount                float32  `json:"orderAmount"`
	RedirectUrl           string   `json:"redirectURI"`
	RedirectURIDeeplink   string   `json:"redirectURIDeeplink"`
	PaymentPageHeaderText string   `json:"paymentPageHeaderText"`
	Narration             string   `json:"narration"`
	AllowedModes          []string `json:"allowedModes"`
}

type CreateTxnRes struct {
	StatusCode string        `json:"statusCode"`
	Message    string        `json:"message"`
	Data       CreateTxnData `json:"data"`
}
type CreateTxnData struct {
	PaymentAmount      float32 `json:"paymentAmount"`
	OrderId            string  `json:"orderId"`
	PaymentURI         string  `json:"paymentURI"`
	PaymentURIDeeplink string  `json:"paymentURIDeeplink"`
}
type CheckTxnReq struct {
	Mid     string `json:"mid"`
	OrderID string `json:"orderId"`
}
type CheckTxnRes struct {
	StatusCode string          `json:"statusCode"`
	Message    string          `json:"message"`
	Data       CheckTxnResData `json:"data"`
}
type CheckTxnResData struct {
	PaymentAmount   float32          `json:"paymentAmount"`
	PaymentStatus   string           `json:"paymentStatus"`
	Currency        string           `json:"currency"`
	OrderId         string           `json:"orderId"`
	PaymentURI      *string          `json:"paymentURI"`
	OrderAmount     *string          `json:"orderAmount"`
	PaymentRefId    *string          `json:"paymentRefId"`
	BeneficiaryName *string          `json:"beneficiaryName"`
	Payment         *CheckTxnPayment `json:"payments"`
}
type CheckTxnPayment struct {
	Amount      float32 `json:"amount"`
	Mode        string  `json:"mode"`
	Status      string  `json:"status"`
	CompletedAt int64   `json:"completedAt"`
	RefId       string  `json:"refId"`
}
