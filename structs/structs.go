package structs

type Account struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
type Transaction struct {
	ID          int64  `json:"id"`
	Sender      string `json:"sender"`
	Beneficiary string `json:"beneficiary"`
	Amount      int64  `json:"amount"`
	Remark      string `json:"remark"`
}
