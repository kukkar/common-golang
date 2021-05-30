package mobilenotification

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kukkar/common-golang/pkg/logger"
)

//ConnectOneAccessor
type ConnectOneAccessor struct {
	IPPort   string
	Version  string
	UserName string
	Password string
}

func (this ConnectOneAccessor) SendMessage(number string, message string) (*SendOTPRes, error) {

	parsedURL := fmt.Sprintf("%s?dest=%s&msg=%s&uname=%s&pwd=%s&send=BHRTPE", this.IPPort, number, url.QueryEscape(message), this.UserName, this.Password)
	var output SendOTPRes
	logger.Logger.Info(parsedURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", parsedURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if 200 != res.StatusCode {
		return nil, fmt.Errorf("send otp to merchant failed with code %v and body %v", res.StatusCode, res.Body)
	}
	output.Status = true
	return &output, nil
}
