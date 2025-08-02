package accountHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/account"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetAccount(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params account.AccountForm
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

		account, err := db_sql.GetAccount(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, account)
	}
}

func GetAccounts(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params account.AccountForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		if params.House_User == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No house_userID provided"})
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

		accounts, err := db_sql.GetAccounts(db, params.House_User)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, accounts)
	}
}

func CreateAccount(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params account.AccountForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		if params.House_User == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No house_userID provided"})
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

		id, err := db_sql.CreateAccount(db, params.Name, params.House_User, params.Amount, params.Currency, params.TheoreticalAmount)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"accountID": id})
	}
}

func UpdateAccount(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params account.AccountForm
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No id provided"})
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

		_, err = db_sql.ChangeAccount(db, params.Id, params.Name, params.Amount, params.TheoreticalAmount)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		account, err := db_sql.GetAccount(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"account": account})
	}
}

//TODO: delete account
