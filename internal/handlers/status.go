package status

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type status struct {
	Database bool `json:"isDatabaseOnline"`
}

func GetDBStatus(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		state := &status{
			Database: true,
		}
		pingErr := db.Ping()
		if pingErr != nil {
			slog.Error(pingErr.Error())
			state.Database = false
		}

		ctx.JSON(http.StatusOK, state)
	}
}
