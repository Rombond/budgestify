package router

import (
	"database/sql"
	"os"

	accountHouseHandlers "github.com/Rombond/budgestify/internal/handlers/accountHouse"
	categoryHandlers "github.com/Rombond/budgestify/internal/handlers/category"
	houseHandlers "github.com/Rombond/budgestify/internal/handlers/house"
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

		houseGroup := api.Group("/houses")
		{
			houseGroup.GET("/one", houseHandlers.GetHouse(db))
			houseGroup.GET("/all", houseHandlers.GetHouses(db))

			houseGroup.POST("/link", houseHandlers.CreateHouseForUser(db))
			houseGroup.POST("/invite", houseHandlers.CreateHouseForUser(db))
			houseGroup.POST("/edit", houseHandlers.ChangeHouse(db))

			accountHouseGroup := houseGroup.Group("/account")
			{
				accountHouseGroup.GET("/", accountHouseHandlers.GetAccountHouse(db))

				accountHouseGroup.POST("/create", accountHouseHandlers.CreateAccountForHouse(db))
				accountHouseGroup.POST("/edit", accountHouseHandlers.UpdateAccountForHouse(db))
			}
		}

		categoryGroup := api.Group("/categories")
		{
			categoryGroup.GET("/one", categoryHandlers.GetCategory(db))
			categoryGroup.GET("/all", categoryHandlers.GetCategories(db))

			categoryGroup.POST("/create", categoryHandlers.CreateCategory(db))
			categoryGroup.POST("/edit", categoryHandlers.UpdateCategory(db))
		}
	}

	router.Run(":" + os.Getenv("API_PORT"))
}
