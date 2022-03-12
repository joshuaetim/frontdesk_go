package infrastructure_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestVisitorSave(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)
	staff := seedStaff(t, db, user.ID)

	visitor := seedVisitor(t, db, user.ID, staff.ID)
	assert.EqualValues(t, 1, visitor.ID)
}

func TestVisitorGet(t *testing.T) {
	initTest(t)
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)
	staff := seedStaff(t, db, user.ID)

	visitor := seedVisitor(t, db, user.ID, staff.ID)

	vr := infrastructure.NewVisitorRepository(db)
	v, err := vr.GetVisitor(visitor.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, visitor.ID, v.ID)
}

func TestVisitorGetAll(t *testing.T) {
	initTest(t)
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)
	staff := seedStaff(t, db, user.ID)

	var visitors []model.Visitor 
	for i := 0; i < 4; i++ {
		visitor := seedVisitor(t, db, user.ID, staff.ID)
		visitors = append(visitors, visitor)
	}

	sr := infrastructure.NewVisitorRepository(db)
	visitorsGet, err := sr.GetAllVisitor()
	assert.Nil(t, err)

	assert.EqualValues(t, len(visitors), len(visitorsGet))
}

func TestVisitorUpdate(t *testing.T) {
	initTest(t)
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)
	staff := seedStaff(t, db, user.ID)

	visitor := seedVisitor(t, db, user.ID, staff.ID)

	vr := infrastructure.NewVisitorRepository(db)
	visitor.Firstname = "Name Updated"
	v2, err := vr.UpdateVisitor(visitor)
	assert.Nil(t, err)

	assert.EqualValues(t, "Name Updated", v2.Firstname)
}

func TestVisitorDelete(t *testing.T) {
	initTest(t)
	initTest(t)
	db := infrastructure.DB()
	user := seedUser(t, db)
	staff := seedStaff(t, db, user.ID)

	visitor := seedVisitor(t, db, user.ID, staff.ID)

	vr := infrastructure.NewVisitorRepository(db)
	err := vr.DeleteVisitor(visitor)
	assert.Nil(t, err)

	_, err = vr.GetVisitor(visitor.ID)
	assert.NotNil(t, err)
}

func seedVisitor(t *testing.T, dbInstance *gorm.DB, user, staff uint) model.Visitor {
	visitor := model.Visitor{
		Firstname: gofakeit.FirstName(),
		Lastname: gofakeit.LastName(),
		Email: gofakeit.Email(),
		Phone: gofakeit.Phone(),
		Occupation: gofakeit.JobTitle(),
		Note: gofakeit.Sentence(30),
		StaffID: staff,
		UserID: user,
	}
	vr := infrastructure.NewVisitorRepository(dbInstance)
	visitor, err := vr.AddVisitor(visitor)
	assert.Nil(t, err)

	return visitor
}