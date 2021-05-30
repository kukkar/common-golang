package mobilenotification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GupShupAccessor
type JapiAccessor struct {
	IPPort   string
	Version  string
	UserName string
	Password string
	AuthKey  string
}

func (this JapiAccessor) SendMessage(number string, message string) (*SendOTPRes, error) {

	//parsedURL := fmt.Sprintf("%s?dest=%s&msg=%s&uname=%s&pwd=%s&send=BHRTPE", this.IPPort, number, url.QueryEscape(merchant), this.UserName, this.Password)
	var output SendOTPRes
	var serviceRes yapiServiceRes
	msgData := make([]msgReq, 0)
	msgData = append(msgData, msgReq{
		Destination: []string{number},
		Text:        message,
		Send:        "BHRTPE",
		Type:        "PM",
	})
	var serviceReq = yapiServiceReq{
		Version:  "1.0",
		Key:      this.AuthKey,
		Encrypt:  "0",
		Messages: msgData,
	}
	url := fmt.Sprintf("%s/%s", this.IPPort, "httpapi/JsonReceiver")
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}

	output = SendOTPRes{
		Status:   true,
		UniqueID: serviceRes.AckID,
	}
	return &output, nil
}
