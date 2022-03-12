package infrastructure_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath = filepath.Dir(b)
)

func TestUserSave(t *testing.T) {
	initTest(t)

	db := infrastructure.DB()

	var user model.User
	user.Firstname = "Joshua Etim"
	user.Lastname = "Etim"
	user.Email = "jetimworks@gmail.com"
	user.Password = "password"
	user.Company = "Jetimworks"

	ur := infrastructure.NewUserRepository(db)

	u, err := ur.AddUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, u.ID)
	assert.EqualValues(t, u.Firstname, "Joshua Etim")
}

func TestUserDuplicateEmail(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

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
	ur := infrastructure.NewUserRepository(db)

	u1, err := ur.AddUser(user1)
	assert.Nil(t, err)
	assert.EqualValues(t, u1.Email, "jetimworks@gmail.com")

	u2, err := ur.AddUser(user2)
	assert.NotNil(t, err)
	assert.EqualValues(t, u2.ID, 0)
}

func TestUserGet(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	u1 := seedUser(t, db)
	ur := infrastructure.NewUserRepository(db)

	u2, err := ur.GetUser(u1.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, u2.Email, u1.Email)
}

func TestUserGetByEmail(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	u1 := seedUser(t, db)
	ur := infrastructure.NewUserRepository(db)

	u2, err := ur.GetByEmail(u1.Email)
	assert.Nil(t, err)
	assert.EqualValues(t, u1.ID, u2.ID)
}

func TestUserGetAll(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	var users []model.User
	for i := 0; i < 4; i++ {
		users = append(users, seedUser(t, db))
	}
	fmt.Println(len(users))

	ur := infrastructure.NewUserRepository(db)
	allUsers, err := ur.GetAllUser()
	assert.Nil(t, err)
	assert.EqualValues(t, len(users), len(allUsers))
}

func TestUserUpdate(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	u := seedUser(t, db)
	assert.EqualValues(t, 1, u.ID)

	u.Email = "changed@gmail.com"
	
	ur := infrastructure.NewUserRepository(db)
	u2, err := ur.UpdateUser(u)
	assert.Nil(t, err)

	assert.EqualValues(t, "changed@gmail.com", u2.Email)
}

func TestUserDelete(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	u := seedUser(t, db)

	ur := infrastructure.NewUserRepository(db)

	err := ur.DeleteUser(u)
	assert.Nil(t, err)

	u, err = ur.GetUser(u.ID)
	assert.NotNil(t, err)
}

func TestUserGetStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()

	u := seedUser(t, db)
	var userStaff []model.Staff

	for i := 0; i < 4; i++ {
		staff := seedStaff(t, db, u.ID)
		userStaff = append(userStaff, staff)
	}

	ur := infrastructure.NewUserRepository(db)
	userStaffGet, err := ur.GetUserStaff(u.ID)
	assert.Nil(t, err)

	assert.EqualValues(t, len(userStaff), len(userStaffGet))
	assert.EqualValues(t, userStaffGet[1].UserID, u.ID)
}

func initTest(t *testing.T) {
	err := godotenv.Load(basepath + "/../" + ".env_test")
	if err != nil {
		t.Fatal(err)
	}
}

func seedUser(t *testing.T, dbInstance *gorm.DB) model.User {
	user := model.User{
		Firstname: gofakeit.FirstName(),
		Lastname: gofakeit.LastName(),
		Email: gofakeit.Email(),
		Password: gofakeit.Word(),
		Company: gofakeit.Company(),
	}
	ur := infrastructure.NewUserRepository(dbInstance)
	u, err := ur.AddUser(user)
	assert.Nil(t, err)

	return u
}