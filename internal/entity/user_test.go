package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T){
	user, err := NewUser("Giovanna", "g@.com", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Giovanna", user.Name)
	assert.Equal(t, "g@.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T){
	user, err := NewUser("Giovanna", "g@.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))

	assert.NotEqual(t, "123456", user.Password)
}