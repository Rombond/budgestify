package transactionHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/transaction"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetTransaction(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params transaction.TransactionForm
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

		transaction, err := db_sql.GetTransaction(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, transaction)
	}
}

func GetTransactionsByHouseUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params transaction.TransactionForm
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

		categories, err := db_sql.GetTransactionsByHouseUser(db, params.House_User)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, categories)
	}
}

func CreateTransaction(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params transaction.TransactionForm
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

		categoryID := &params.CategoryID
		if *categoryID == 0 {
			categoryID = nil
		}

		payerAccountID := &params.PayerAccountID
		if *payerAccountID == 0 {
			payerAccountID = nil
		}

		id, err := db_sql.CreateTransaction(db, params.Name, categoryID, params.Amount, params.PayerID, payerAccountID, params.PayDate, params.Currency, params.ConversionRate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"transactionID": id})
	}
}

func UpdateTransaction(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params transaction.TransactionForm
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

		categoryID := &params.CategoryID
		if *categoryID == 0 {
			categoryID = nil
		}

		payerAccountID := &params.PayerAccountID
		if *payerAccountID == 0 {
			payerAccountID = nil
		}

		_, err = db_sql.ChangeTransaction(db, params.Id, params.Name, categoryID, params.Amount, params.PayerID, payerAccountID, params.PayDate, params.Currency, params.ConversionRate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		transaction, err := db_sql.GetTransaction(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

//TODO: delete transaction
