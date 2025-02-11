package statusHandlers

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"

	db_sql "github.com/Rombond/budgestify/internal/sql"

	"github.com/gin-gonic/gin"
)

type status struct {
	Database bool `json:"isDatabaseOnline"`
	Setup    bool `json:"isSetupComplete"`
}

type responseStatus struct {
	Status status `json:"status"`
}

type responseSetup struct {
	Setup db_sql.StateSetup `json:"setup"`
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
		state.Setup = db_sql.GetSetupDone()

		resp := &responseStatus{
			Status: *state,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func GetSetupStatus(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if id > 0 {
			db_sql.UpdateStateSetup(db, id)
		}

		resp := &responseSetup{
			Setup: *db_sql.GetStateSetup(),
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
