package loyalty

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

type GetLoyaltyPointsRes struct {
	Points int
}

type RedeemPointsReq struct {
	TxnID      string
	TxnType    string
	MerchantID int
	Category   string
	Points     string
}

type serviceReqGetLoyalPoints struct {
	MerchantID string `json:"merchant_id"`
}

type serviceResGetLoyaltyPoints struct {
	Points int  `json:"points"`
	Status bool `json:"status"`
}

type loyaltyServiceErrRes struct {
	Status bool `json:"status"`
}

type serviceReqRedeemPoints struct {
	TxnID      string `json:"txn_id"`
	TxnType    string `json:"txn_type"`
	MerchantID string `json:"merchant_id"`
	Category   string `json:"category"`
	Points     string `json:"points"`
}

type serviceResRedeemPoints struct {
	Status bool `json:"status"`
}
