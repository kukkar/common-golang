package lending

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	clientvalidator clientvalidator.ClientValidator
}

func (this Accessor) AllowBankAccountChange(clientName string, merchantID int) (bool, error) {

	var serviceRes allowBankAccountChangeRes
	hmacReq := make(map[string]interface{})
	url := fmt.Sprintf("%s/%s/%s?merchantId=%d", this.IPPort, lendingRoute, allowedBankChangeRoute, merchantID)

	hmacReq["merchantId"] = strconv.Itoa(merchantID)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return false, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(PennyDropHeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(PennyDropHeaderClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if 200 != resp.StatusCode {
		return false, fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return false, err
	}
	if !serviceRes.Success {
		return false, fmt.Errorf(fmt.Sprintf("%s", body))
	}
	return serviceRes.Data.BankAccountChange, nil
}
