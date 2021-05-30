package mobilenotification

import (
	"fmt"
	"time"
)

//Config
type Config struct {
	ServiceToUse string
	IPPort       string
	UserName     string
	Password     string
	AuthKey      string
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" || this.UserName == "" || this.Password == "" {
		return fmt.Errorf("invalid config url username password required")
	}
	return nil
}

type yourMessageRes []struct {
	Code              string   `json:"code"`
	Desc              string   `json:"desc"`
	ReqID             string   `json:"reqId"`
	Time              string   `json:"time"`
	PartMessageIds    []string `json:"partMessageIds"`
	TotalMessageParts int      `json:"totalMessageParts"`
}
type yourMsgResData struct {
	Code string `json:"CODE"`
	Info string `json:"INFO"`
}

type SendOTPRes struct {
	Status   bool
	Resposne string
	UniqueID string
}

type yapiServiceReq struct {
	Version  string   `json:"ver"`
	Key      string   `json:"key"`
	Encrypt  string   `json:"encrypt"`
	Messages []msgReq `json:"messages"`
}

type msgReq struct {
	Destination []string `json:"dest"`
	Text        string   `json:"Text"`
	Send        string   `json:"send"`
	Type        string   `json:"type"`
}

type yapiServiceRes struct {
	AckID  string             `json:"ackid"`
	Time   time.Time          `json:"time"`
	Status yapiServiceResData `json:"status"`
}
type yapiServiceResData struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}
