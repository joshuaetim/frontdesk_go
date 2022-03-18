package handler

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonReq struct {
	Name string `json:"name"`
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func JSONRequest(ctx *gin.Context) {
	var req JsonReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request: "+err.Error())
		return
	}
	// fmt.Println(req.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"data": req.Name,
	})
}

func CheckAuth(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	ctx.JSON(http.StatusOK, userID)
}
