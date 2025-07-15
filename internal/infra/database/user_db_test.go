package database

import (
	"testing"

	"goexpert-api/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

  if err != nil {
    t.Error(err)
  }

  db.AutoMigrate(&entity.User{})
  user, _ := entity.NewUser("tone", "tone@tone.com", "123456")
  userDB := NewUser(db)

  err = userDB.Create(user)
  if err != nil {
    t.Error(err)
  }

  assert.Nil(t, err)

  var bdUser entity.User
  err = db.First(&bdUser, "id = ?",user.ID).Error
  assert.Nil(t, err)
  assert.Equal(t, user.Email, bdUser.Email)
  assert.Equal(t, user.Name, bdUser.Name)
  assert.Equal(t, user.ID, bdUser.ID)
  assert.NotNil(t,  bdUser.Password)
}


func TestFindByEmail(t *testing.T){
  bd, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

  if err != nil {
    t.Error(err)
  }

  bd.AutoMigrate(&entity.User{})

  user, _ := entity.NewUser("cenas", "cenas@cenas.pt", "1234")
  userDB := NewUser(bd)

  err = userDB.Create(user)

  if err != nil {
    t.Error(err)
  }


  userFind, err := userDB.FindByEmail(user.Email)

  if err != nil {
    t.Error(err)
  }

  assert.Nil(t, err)
  assert.Equal(t, user.Email, userFind.Email)
  assert.Equal(t, user.Name, userFind.Name)
  assert.Equal(t, user.ID, userFind.ID)
  assert.NotNil(t,  userFind.Password)
}

