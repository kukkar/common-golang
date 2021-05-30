package pennydrop

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//PennyDropAccessor
type PennyDropAccessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	clientvalidator clientvalidator.ClientValidator
}

func (this PennyDropAccessor) PennyDrop(request PennyDropReq) (*AccountInfo, error) {

	var serviceRes pennyDropRes
	var output AccountInfo
	var serviceReq = pennyDropServiceReq{
		IFSC:          request.IFSC,
		AccountNumber: request.AccountNumber,
	}
	url := fmt.Sprintf("%s%s/%s", this.IPPort, routePennyDrop, this.Version)
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}
	var hmacReq map[string]interface{}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return nil, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateHMac(c, hmacReq,
		request.ClientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add(globalconst.HMAC, fmt.Sprintf("%x", hmac))
	req.Header.Add(globalconst.CLIENT_NAME, request.ClientName)

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
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&output, serviceRes.Data)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
