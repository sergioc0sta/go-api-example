package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T){
  newUser, err := NewUser("tone", "tone@tone.pt", "12345")
  assert.Nil(t, err)
  assert.NotNil(t, newUser)
  assert.NotEmpty(t, newUser.ID)
  assert.NotEmpty(t, newUser.Password)
  assert.Equal(t, "tone", newUser.Name)
  assert.Equal(t, "tone@tone.pt", newUser.Email)
}


func TestPasswordNewUser(t *testing.T){
  newUser, err := NewUser("tone", "tone@tone.pt", "12345")
  assert.Nil(t, err)
  assert.True(t, newUser.ValidatePassword("12345"))
  assert.False(t, newUser.ValidatePassword("123456"))
  assert.NotEqual(t, "123456", newUser.Password)
}
