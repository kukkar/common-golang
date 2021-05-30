package otpLess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	ClientId        string
	ClientSecret    string
	clientvalidator clientvalidator.ClientValidator
}

func (this Accessor) CreateIntent() (*CreateIntentRes, error) {
	var serviceRes CreateIntentRes

	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/user/getSignupUrl")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(HeaderClientId, this.ClientId)
	req.Header.Add(HeaderClientSecret, this.ClientSecret)
	fmt.Printf("http Req %v url %v", req, url)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if 200 != res.StatusCode {
		return nil, fmt.Errorf(fmt.Sprintf("body %s status code %v", body, res.StatusCode))
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	logger.Info("Verify token Res:", serviceRes)
	if serviceRes.Status != "SUCCESS" {
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &serviceRes, nil
}

func (this Accessor) VerifyWTToken(token string) (*VerifyWTTokenRes, error) {

	var serviceRes VerifyWTTokenRes

	//serviceRes.Mobile = "917069914791"
	//return &serviceRes, nil

	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/user/getUserDetails")
	var otplessServiceReq VerifyTokenOtplLessReq

	otplessServiceReq.Token = token

	serviceReq, err := json.Marshal(otplessServiceReq)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(serviceReq))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(HeaderClientId, this.ClientId)
	req.Header.Add(HeaderClientSecret, this.ClientSecret)
	fmt.Printf("http Req %v url %v", req, url)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if 200 != res.StatusCode {
		return nil, fmt.Errorf(fmt.Sprintf("body %s status code %v", body, res.StatusCode))
	}
	fmt.Printf("http Res %v", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	logger.Info("Verify token Res:", serviceRes)
	if serviceRes.Status != "SUCCESS" {
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &serviceRes, nil
}
