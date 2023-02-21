package structs

type Account struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
type Transaction struct {
	ID          string `json:"id"`
	Sender      string `json:"sender"`
	Beneficiary string `json:"beneficiary"`
	Remark      string `json:"Remark"`
}
