package process_transaction

type TransactionDtoInput struct {
	ID	string 				`json:"id,omitempty"`
	AccountID string 		`json:"account_id,omitempty"`
	Amount float64 			`json:"amount,omitempty"`
}

type TransactionDtoOutPut struct {
	ID string				`json:"id,omitempty"`
	Status string 			`json:"status,omitempty"`
	ErrorMessage string 	`json:"error_message,omitempty"`
}
