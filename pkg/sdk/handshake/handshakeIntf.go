package handshake

type SDK interface {
	PennyDrop(req PennyDropReq) (*PennyDropRes, error)
	GetBeneficiary(req GetBeneficiaryReq) (*GetBeneficiaryResponse, error)
	GetCategoryTree() (*GetCategoryTree, error)
}
