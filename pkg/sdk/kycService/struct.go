package kycService

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

type GetPanHolderNameReq struct {
	PanNumber  string `json:"panNumber"`
	MerchantId int    `json:"identifier"`
	UserType   string `json:"userType"`
}

type GetPanHolderNameRes struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    GetPanHolderNameData `json:"data"`
}

type GetPanHolderNameData struct {
	Name string `json:"name"`
}

type UpdateKYCReq struct {
	DocType   string `json:"docType"`
	PanNumber string `json:"panNumber"`
	Source    string `json:"source"`
}
type UpdateKYCRes struct {
	Status string `json:"status"`
}

type NameMatchReq struct {
	Name1 string `json:"name1"`
	Name2 string `json:"name2"`
}
type NameMatchRes struct {
	Status              bool    `json:"status"`
	NameMatch           bool    `json:"nameMatch"`
	NameMatchPercentage float64 `json:"nameMatchPer"`
}

type GetPanReq struct {
	MerchantId int    `json:"identifier"`
	UserType   string `json:"userType"`
}

type GetPanRes struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    GetPanData `json:"data"`
}

type GetPanData struct {
	PanNo string `json:"panNo"`
	Name  string `json:"name"`
}
