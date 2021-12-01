package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T ) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000

	errValidation := transaction.IsValid()

	assert.Error(t, errValidation)
	assert.Equal(t, "you dont have limit for this transaction", errValidation.Error())
}

func TestTransactionWithAmountLesserThan1(t *testing.T ) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	errValidation := transaction.IsValid()

	assert.Error(t, errValidation)
	assert.Equal(t, "the amount must be greater than 1", errValidation.Error())
}