package router

import (
	"database/sql"
	"os"

	status "github.com/Rombond/budgestify/internal/handlers/status"
	userHandlers "github.com/Rombond/budgestify/internal/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) {
	router := gin.Default()

	router.GET("/", status.GetDBStatus(db))

	api := router.Group("/api")
	{
		api.GET("/status", status.GetDBStatus(db))
		api.GET("/setup", status.GetSetupStatus())

		api.POST("/login", userHandlers.LoginUser(db))

		userGroup := api.Group("/users")
		{
			userGroup.GET("/:id", userHandlers.GetUser(db))

			userGroup.POST("/register", userHandlers.CreateUser(db))
			userGroup.POST("/edit", userHandlers.ChangeUser(db))
		}
	}

	router.Run(":" + os.Getenv("API_PORT"))
}
