package pushnotification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils"
	"go.elastic.co/apm"
)

var _ SDK = (*notificationAccessor)(nil)

//SendGridAccessor
type notificationAccessor struct {
	ipPort  string
	txn     *apm.Transaction
	ctx     context.Context
	rc      *utils.RequestContext
	authKey string
}

//
// Exposed:
// GetProfiled service for usiong with Elastic apm.
//
func GetProfiledService(ctx context.Context, rc *utils.RequestContext, txn *apm.Transaction,
	iPPort string, authKey string) (SDK, error) {
	if iPPort == "" {
		return nil, fmt.Errorf("ipPort can not be empty")
	}

	return &notificationAccessor{
		iPPort,
		txn,
		ctx,
		rc,
		authKey,
	}, nil
}

func (this *notificationAccessor) SendPushNotification(req RequestNotification) error {
	if this.ctx == nil &&
		this.txn != nil {
		this.txn.StartSpan("SendPushNotification", "custom", nil)
	} else {
		if this.ctx != nil {
			span, _ := apm.StartSpan(this.ctx, "SendPushNotification", "custom")
			defer span.End()
		}
	}

	var serviceRes pnServiceRes
	var serviceReq = pnServiceReq{
		To: req.To,
		Data: pnServiceReqData{
			Title:     req.Title,
			Body:      req.Body,
			SoundName: req.SoundName,
			Image:     req.Image,
			ImageType: req.ImageType,
			URL:       req.URL,
		},
	}
	url := this.ipPort
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Authorization", this.authKey)
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("##### service res %v", serviceRes))
	return nil
}
