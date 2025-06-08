package router

import (
	"database/sql"
	"os"

	accountHandlers "github.com/Rombond/budgestify/internal/handlers/account"
	categoryHandlers "github.com/Rombond/budgestify/internal/handlers/category"
	houseHandlers "github.com/Rombond/budgestify/internal/handlers/house"
	recurrenceHandlers "github.com/Rombond/budgestify/internal/handlers/recurrence"
	status "github.com/Rombond/budgestify/internal/handlers/status"
	transactionHandlers "github.com/Rombond/budgestify/internal/handlers/transaction"
	userHandlers "github.com/Rombond/budgestify/internal/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) {
	router := gin.Default()

	router.GET("/", status.GetDBStatus(db))

	api := router.Group("/api")
	{
		api.GET("/status", status.GetDBStatus(db))
		api.GET("/setup", status.GetSetupStatus(db))
		api.GET("/setup/:id", status.GetSetupStatus(db))

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
		}

		categoryGroup := api.Group("/categories")
		{
			categoryGroup.GET("/one", categoryHandlers.GetCategory(db))
			categoryGroup.GET("/all", categoryHandlers.GetCategories(db))

			categoryGroup.POST("/create", categoryHandlers.CreateCategory(db))
			categoryGroup.POST("/edit", categoryHandlers.UpdateCategory(db))
		}

		accountGroup := api.Group("/accounts")
		{
			accountGroup.GET("/one", accountHandlers.GetAccount(db))
			accountGroup.GET("/all", accountHandlers.GetAccounts(db))

			accountGroup.POST("/create", accountHandlers.CreateAccount(db))
			accountGroup.POST("/edit", accountHandlers.UpdateAccount(db))
		}

		transactionGroup := api.Group("/transactions")
		{
			transactionGroup.GET("/one", transactionHandlers.GetTransaction(db))
			transactionGroup.GET("/all", transactionHandlers.GetTransactionsByHouseUser(db))

			transactionGroup.POST("/create", transactionHandlers.CreateTransaction(db))
			transactionGroup.POST("/edit", transactionHandlers.UpdateTransaction(db))
		}

		recurrenceGroup := api.Group("/recurrences")
		{
			recurrenceGroup.GET("/one", recurrenceHandlers.GetRecurrence(db))
			recurrenceGroup.GET("/all", recurrenceHandlers.GetRecurrencesByHouseUser(db))

			recurrenceGroup.POST("/create", recurrenceHandlers.CreateRecurrence(db))
			recurrenceGroup.POST("/edit", recurrenceHandlers.UpdateRecurrence(db))
		}
	}

	router.Run(":" + os.Getenv("API_PORT"))
}
