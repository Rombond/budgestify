package router

import (
	"database/sql"
	"os"

	status "github.com/Rombond/budgestify/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) {
	router := gin.Default()

	router.GET("/", status.GetDBStatus(db))

	router.Run(":" + os.Getenv("API_PORT"))
}
