package pennydrop

type Sdk interface {
	PennyDrop(req PennyDropReq) (*AccountInfo, error)
}
