package infrastructure_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSaveStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	u := seedUser(t, db)

	staff := seedStaff(t, db, int(u.ID))
	assert.EqualValues(t, 1, staff.ID)
}

func TestGetStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)

	staff := seedStaff(t, db, int(user.ID))

	sr := infrastructure.NewStaffRepository(db)
	staffGet, err := sr.GetStaff(int(staff.ID))
	assert.Nil(t, err)

	assert.EqualValues(t, staff.ID, staffGet.ID)
}

func TestGetAllStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)

	var allStaff []model.Staff

	for i := 1; i <= 4; i++ {
		staff := seedStaff(t, db, int(user.ID))
		allStaff = append(allStaff, staff)
	}

	sr := infrastructure.NewStaffRepository(db)
	allStaffGet, err := sr.GetAllStaff()
	assert.Nil(t, err)

	assert.EqualValues(t, len(allStaff), len(allStaffGet))
}

func TestUpdateStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)

	staff := seedStaff(t, db, int(user.ID))

	sr := infrastructure.NewStaffRepository(db)
	s1, err := sr.GetStaff(int(staff.ID))
	assert.Nil(t, err)

	s1.Firstname = "Johnny Update"

	s2, err := sr.UpdateStaff(s1)
	assert.Nil(t, err)

	assert.EqualValues(t, "Johnny Update", s2.Firstname)
}

func TestDeleteStaff(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)

	staff := seedStaff(t, db, int(user.ID))

	sr := infrastructure.NewStaffRepository(db)
	err := sr.DeleteStaff(staff)
	assert.Nil(t, err)

	_, err = sr.GetStaff(int(staff.ID))
	assert.NotNil(t, err)
}

func seedStaff(t *testing.T, dbInstance *gorm.DB, user int) model.Staff {
	staff := model.Staff{
		Firstname: gofakeit.FirstName(),
		Lastname: gofakeit.LastName(),
		Email: gofakeit.Email(),
		Phone: gofakeit.Phone(),
		Address: gofakeit.Address().Address,
		Status: "active",
		UserID: user,
	}
	sr := infrastructure.NewStaffRepository(dbInstance)
	s, err := sr.AddStaff(staff)
	assert.Nil(t, err)

	return s
}