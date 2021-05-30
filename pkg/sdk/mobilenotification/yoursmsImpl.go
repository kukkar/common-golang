package mobilenotification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kukkar/common-golang/pkg/logger"
)

//YourMessageAccessor
type YourMessageAccessor struct {
	IPPort   string
	Version  string
	UserName string
	AuthKey  string
}

func (this YourMessageAccessor) SendMessage(number string, message string) (*SendOTPRes, error) {

	parsedURL := fmt.Sprintf("%s?username=%s&apikey=%s&signature=%s&dest=%s&msgtxt=%s&msgtype=PM",
		this.IPPort, this.UserName, this.AuthKey, YourSMSSender, number, url.QueryEscape(message))
	var finalRes SendOTPRes
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var output yourMessageRes
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	if 200 != res.StatusCode {
		return nil, fmt.Errorf("send otp to merchant failed with code %v and body %v", res.StatusCode, res.Body)
	}
	if len(output) > 0 {
		if output[0].ReqID == "" {
			return nil, fmt.Errorf("send otp to merchant failed Error :  %v", output)
		}
	}
	finalRes.Status = true
	return &finalRes, nil
}
