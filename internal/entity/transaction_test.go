package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Joana Doe", "ja@j.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 500)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1, transaction.AccountFrom)
	assert.Equal(t, account2, transaction.AccountTo)

	assert.Equal(t, 500.0, account1.Balance)
	assert.Equal(t, 1500.0, account2.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Joana Doe", "ja@j.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 1500)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient balance", err.Error())
}
