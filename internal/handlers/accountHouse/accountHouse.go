package accountHouseHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/accountHouse"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetAccountHouse(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params accountHouse.AccountHouse
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		valid, err := token.IsTokenValid(ctx.GetHeader("Authorization"), params.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Authorization is not valid"})
			return
		}

		house, err := db_sql.GetAccountHouse(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		house.UserId = params.UserId
		ctx.JSON(http.StatusOK, house)
	}
}

func CreateAccountForHouse(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params accountHouse.AccountHouseForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}
		if params.HouseID == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No houseID provided"})
			return
		}
		if params.Name == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No name provided"})
			return
		}

		valid, err := token.IsTokenValid(ctx.GetHeader("Authorization"), params.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Authorization is not valid"})
			return
		}

		isAdmin, err := db_sql.IsUserAdmin(db, params.HouseID, params.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !isAdmin {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "You are not an administrator of the house"})
			return
		}

		id, err := db_sql.CreateAccountHouse(db, params.Name, params.Amount)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		_, err = db_sql.AddAccount(db, params.HouseID, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			db_sql.DeleteAccountHouse(db, id)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"accountHouseID": id})
	}
}

func UpdateAccountForHouse(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params accountHouse.AccountHouse
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		if params.Id == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No id found"})
			return
		}

		valid, err := token.IsTokenValid(ctx.GetHeader("Authorization"), params.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Authorization is not valid"})
			return
		}

		canEdit, err := db_sql.DoesUserCanEditAccount(db, params.Id, params.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !canEdit {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "You are not an administrator of the house"})
			return
		}

		if params.Amount != 0 {
			_, err = db_sql.ChangeAccountHouseAmount(db, params.Id, params.Amount)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		if params.Name != "" {
			_, err = db_sql.ChangeAccountHouseName(db, params.Id, params.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		accountHouse, err := db_sql.GetAccountHouse(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, accountHouse)
	}
}

//TODO: delete accountHouse
