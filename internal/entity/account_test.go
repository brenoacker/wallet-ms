package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, 0.0, account.Balance)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateNewAccountWhenClientIsNil(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	err := account.Debit(50)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, account.Balance)
}

func TestDebitAccountWhenBalanceIsInsufficient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := account.Debit(50)
	assert.NotNil(t, err)
	assert.Equal(t, "insufficient balance", err.Error())
}
