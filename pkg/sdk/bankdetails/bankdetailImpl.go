package bankdetail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kukkar/common-golang/pkg/logger"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
}

func (this Accessor) SubmitBankDetail(bankDetailReq SubmitBankDetail) (*SubmitBankDetailRes, error) {

	var res SubmitBankDetailRes
	url := fmt.Sprintf("%s%s/%s", this.IPPort, this.Version, SubmitBankDetailRoute)
	logger.Logger.Info(fmt.Sprintf("submit bank detail request %v", bankDetailReq))
	j, err := json.Marshal(bankDetailReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
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
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (this Accessor) GetBankDetail(bankDetailReq GetBankDetailRequest) (*GetBankDetailRes, error) {

	var res GetBankDetailRes
	url := fmt.Sprintf("%s/%s", this.IPPort, GetBankDetailRoute)
	j, err := json.Marshal(bankDetailReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
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
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
