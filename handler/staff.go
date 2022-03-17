package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/domain/repository"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"gorm.io/gorm"
)

type StaffHandler interface {
	GetStaff(*gin.Context)
	CreateStaff(*gin.Context)
	GetAllStaff(*gin.Context)
	UpdateStaff(*gin.Context)
	DeleteStaff(*gin.Context)
}

type staffHandler struct {
	repo repository.StaffRepository
}

func NewStaffHandler(db *gorm.DB) StaffHandler {
	return &staffHandler{
		repo: infrastructure.NewStaffRepository(db),
	}
}

func (sh *staffHandler) CreateStaff(ctx *gin.Context) {
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	// TODO: check for empty fields
	staff, err := sh.repo.AddStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": staff})
}

func (sh *staffHandler) GetStaff(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "value error: " + err.Error()})
		return
	}
	staff, err := sh.repo.GetStaff(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble fetching staff: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff})
}

func (sh *staffHandler) GetAllStaff(ctx *gin.Context) {
	staff, err := sh.repo.GetAllStaff()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem fetching user; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff})
}

func (sh *staffHandler) UpdateStaff(ctx *gin.Context) {
	staffID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem with input data: " + err.Error()})
		return
	}
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staff.ID = uint(staffID)
	staff, err = sh.repo.UpdateStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem updating user; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff})
}

func (sh *staffHandler) DeleteStaff(ctx *gin.Context) {
	staffID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem with input data: " + err.Error()})
		return
	}
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staff.ID = uint(staffID)
	err = sh.repo.DeleteStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem updating user; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "staff deleted"})
}
