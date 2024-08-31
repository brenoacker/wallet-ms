package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	name := "John Doe"
	email := "j@j.com"
	client, err := NewClient(name, email)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	name := ""
	email := ""
	client, err := NewClient(name, email)
	assert.NotNil(t, err)
	assert.Nil(t, client)
}
func TestUpdateClient(t *testing.T) {
	name := "John Doe"
	email := "j@j.com"
	client, _ := NewClient(name, email)
	newName := "Jane Doe"
	newEmail := "jdoe@j.com"
	err := client.Update(newName, newEmail)
	assert.Nil(t, err)
	assert.Equal(t, newName, client.Name)
	assert.Equal(t, newEmail, client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	name := "John Doe"
	email := "j@j.com"
	client, _ := NewClient(name, email)
	newName := ""
	newEmail := ""
	err := client.Update(newName, newEmail)
	assert.NotNil(t, err)
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
