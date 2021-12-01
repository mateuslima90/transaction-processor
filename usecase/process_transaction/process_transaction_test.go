package process_transaction

import (
	"github.com/golang/mock/gomock"
	mock_entity "github.com/mateuslima90/transaction-processor/entity/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := TransactionDtoInput {
		ID: "1",
		AccountID: "1",
		Amount: 200,
	}

	expectedOuput := TransactionDtoOutPut {
		ID: "1",
		Status: "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_entity.NewMockTransactionRepository(ctrl)
	repository.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)


	usecase := NewProcessTransaction(repository)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOuput, output)
}
