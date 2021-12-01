package process_transaction

import "github.com/mateuslima90/transaction-processor/entity"

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutPut, error) {
	transation := entity.NewTransaction()
	transation.ID = input.ID
	transation.AccountID = input.AccountID
	transation.Amount = input.Amount
	invalidTransaction := transation.IsValid()

	if invalidTransaction != nil {
		return p.rejectTransaction(transation, invalidTransaction)
	}
	return p.approvedTransaction(transation)
}

func (p *ProcessTransaction) approvedTransaction(transation *entity.Transaction) (TransactionDtoOutPut, error) {
	err := p.Repository.Insert(transation.ID, transation.AccountID, transation.Amount, "approved", "")
	if err != nil {
		return TransactionDtoOutPut{}, err
	}
	output := TransactionDtoOutPut{
		ID:           transation.ID,
		Status:       "approved",
		ErrorMessage: "",
	}
	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transation *entity.Transaction, invalidTransaction error) (TransactionDtoOutPut, error) {
	err := p.Repository.Insert(transation.ID, transation.AccountID, transation.Amount, "rejected", invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutPut{}, err
	}
	output := TransactionDtoOutPut{
		ID:           transation.ID,
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}
	return output, nil
}