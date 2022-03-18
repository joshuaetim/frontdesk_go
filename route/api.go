package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/frontdesk/handler"
	"github.com/joshuaetim/frontdesk/infrastructure"
	"github.com/joshuaetim/frontdesk/middleware"
)

func RunAPI(address string) error {
	db := infrastructure.DB()
	userHandler := handler.NewUserHandler(db)
	staffHandler := handler.NewStaffHandler(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Frontdesk v1")
	})
	apiRoutes := r.Group("/api")

	apiRoutes.GET("/checkauth", middleware.AuthorizeJWT(), handler.CheckAuth)

	userRoutes := apiRoutes.Group("/auth")
	userRoutes.POST("/register", userHandler.CreateUser)
	userRoutes.POST("/login", userHandler.SignInUser)

	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeJWT())
	userProtectedRoutes.GET("/:id", userHandler.GetUser)

	staffRoutes := apiRoutes.Group("/staff", middleware.AuthorizeJWT())
	staffRoutes.GET("/:id", staffHandler.GetStaff)
	staffRoutes.POST("/", staffHandler.CreateStaff)
	staffRoutes.GET("/", staffHandler.GetAllStaffByUser)

	return r.Run(address)
}
