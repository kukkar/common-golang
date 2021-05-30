package pushnotification

type RequestNotification struct {
	To        string
	Title     string
	Body      string
	SoundName string
	Image     string
	ImageType string
	URL       string
}

type pnServiceReq struct {
	To   string           `json:"to"`
	Data pnServiceReqData `json:"data"`
}

type pnServiceReqData struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	SoundName string `json:"soundname"`
	Image     string `json:"image"`
	ImageType string `json:"image-type"`
	URL       string `json:"url"`
}

type pnServiceRes struct {
	MulticastID  int64 `json:"multicast_id"`
	Success      int   `json:"success"`
	Failure      int   `json:"failure"`
	CanonicalIds int   `json:"canonical_ids"`
	Results      []struct {
		MessageID string `json:"message_id"`
	} `json:"results"`
}
