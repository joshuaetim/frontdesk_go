package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/frontdesk/handler"
	"github.com/joshuaetim/frontdesk/infrastructure"
)

func RunAPI(address string) error {
	db := infrastructure.DB()
	userHandler := handler.NewUserHandler(db)
	// staffHandler := handler.NewStaffHandler(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Frontdesk v1")
	})
	apiRoutes := r.Group("/api")

	userRoutes := apiRoutes.Group("/user")
	userRoutes.POST("/register", userHandler.CreateUser)
	userRoutes.POST("/login", userHandler.SignInUser)

	return r.Run(address)
}
