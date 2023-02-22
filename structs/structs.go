package structs

type Account struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Balance int64  `json:"balance,omitempty"`
}
type Transaction struct {
	ID          int64  `json:"id,omitempty"`
	Sender      int64  `json:"sender,omitempty"`
	Beneficiary int64  `json:"beneficiary,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type Record struct {
	ID              int64  `json:"id"`
	SenderID        int64  `json:"sender_id"`
	SenderName      string `json:"sender_name"`
	BeneficiaryID   int64  `json:"beneficiary_id"`
	BeneficiaryName string `json:"beneficiary_name"`
	Amount          string `json:"amount"`
	Remark          string `json:"remark"`
}
