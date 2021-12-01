package process_transaction

type TransactionDtoInput struct {
	ID	string
	AccountID string
	Amount float64
}

type TransactionDtoOutPut struct {
	ID string
	Status string
	ErrorMessage string
}
