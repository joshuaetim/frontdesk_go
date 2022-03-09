package infrastructure_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	initTest(t)

	var user model.User
	user.Firstname = "Joshua Etim"
	user.Lastname = "Etim"
	user.Email = "jetimworks@gmail.com"
	user.Password = "password"
	user.Company = "Jetimworks"

	ur := infrastructure.NewUserRepository()

	u, err := ur.AddUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, u.ID)
	assert.EqualValues(t, u.Firstname, "Joshua Etim")
}

func TestDuplicateEmail(t *testing.T) {
	initTest(t)
	user1 := model.User{
		Firstname: "Josh",
		Lastname: "Etim",
		Company: "Jetimworks",
		Email: "jetimworks@gmail.com",
		Password: "password",
	}
	user2 := model.User{
		Firstname: "Josh",
		Lastname: "Etim",
		Company: "Jetimworks",
		Email: "jetimworks@gmail.com",
		Password: "password",
	}
	ur := infrastructure.NewUserRepository()

	u1, err := ur.AddUser(user1)
	assert.Nil(t, err)
	assert.EqualValues(t, u1.Email, "jetimworks@gmail.com")

	u2, err := ur.AddUser(user2)
	assert.NotNil(t, err)
	assert.EqualValues(t, u2.ID, 0)
}

func TestGetUser(t *testing.T) {
	initTest(t)
	user1 := model.User{
		Firstname: "Josh",
		Lastname: "Etim",
		Company: "Jetimworks",
		Email: "jetimworks@gmail.com",
		Password: "password",
	}
	ur := infrastructure.NewUserRepository()
	u1, err := ur.AddUser(user1)
	assert.Nil(t, err)

	u2, err := ur.GetUser(int(u1.ID))
	assert.Nil(t, err)
	assert.EqualValues(t, u2.Email, u1.Email)
}

func TestGetUserByEmail(t *testing.T) {
	initTest(t)
	user1 := model.User{
		Firstname: "Josh",
		Lastname: "Etim",
		Company: "Jetimworks",
		Email: "jetimworks@gmail.com",
		Password: "password",
	}
	ur := infrastructure.NewUserRepository()
	u1, err := ur.AddUser(user1)
	assert.Nil(t, err)

	u2, err := ur.GetByEmail(u1.Email)
	assert.Nil(t, err)
	assert.EqualValues(t, u1.ID, u2.ID)
}

func initTest(t *testing.T) {
	err := godotenv.Load("../.env_test")
	if err != nil {
		t.Fatal(err)
	}
}