package kyc

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

type gstDetailServiceRes struct {
	Success    bool           `json:"success"`
	StatusCode string         `json:"status_code"`
	Message    string         `json:"message"`
	Data       GSTServiceData `json:"data"`
}
type GSTServiceData struct {
	GSTNumber string `json:"gst_number"`
	Status    string `json:"status"`
}

type GSTDetail struct {
	GSTNumber string
	Status    string
}

type PANDetail struct {
	PANNumber string
	Status    string
}

type panServiceRes struct {
	Success    bool           `json:"success"`
	StatusCode string         `json:"status_code"`
	Message    string         `json:"message"`
	Data       PANServiceData `json:"data"`
}
type PANServiceData struct {
	DocumentNumber string `json:"document_number"`
	Status         string `json:"status"`
	Name           string `json:"name"`
}

type PanFetchRes struct {
	Success    bool            `json:"success"`
	StatusCode string          `json:"status_code"`
	Message    string          `json:"message"`
	Data       PanFetchDataRes `json:"data"`
}

type PanFetchDataRes struct {
	NameMatchPercent  float64 `json:"name_match_percent"`
	PanCardHolderName string  `json:"pan_card_holder_name"`
	ProofType         string  `json:"proof_type"`
}

type PanDataVerfiy struct {
	NameMatchPercent  float64 `json:"name_match_percent"`
	PanCardHolderName string  `json:"pan_card_holder_name"`
	ProofType         string  `json:"proof_type"`
}

type PanFetchDataReq struct {
	PanNumber string `json:"pan_card_number"`
	PanName   string `json:"name_to_match"`
}

type SubmitKycDocReq struct {
	ProofNumber     string
	MerchantID      int
	DocType         string
	MerchantStoreID *int
}

type submitKycDocServiceReq struct {
	ProofNumber     string `json:"proof_number"`
	DocType         string `json:"app_name"`
	MerchantID      int    `json:"merchant_id"`
	MerchantStoreID *int   `json:"merchant_store_id"`
}

type submitKycDocServiceRes struct {
	Success    bool   `json:"success"`
	StatusCode string `json:"status_code"`
	Msg        string `json:"msg"`
}
