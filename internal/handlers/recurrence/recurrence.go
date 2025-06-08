package recurrenceHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/recurrence"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetRecurrence(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params recurrence.RecurrenceForm
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

		recurrence, err := db_sql.GetRecurrence(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, recurrence)
	}
}

func GetRecurrencesByHouseUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params recurrence.RecurrenceForm
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

		categories, err := db_sql.GetRecurrencesByHouseUser(db, params.House_User)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, categories)
	}
}

func CreateRecurrence(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params recurrence.RecurrenceForm
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

		categoryID := &params.Category
		if *categoryID == 0 {
			categoryID = nil
		}

		id, err := db_sql.CreateRecurrence(db, params.Name, params.House_User, params.PayerAccountID, categoryID, params.Amount, params.Currency, params.ConversionRate, params.PayDate, params.DayCycle)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"recurrenceID": id})
	}
}

func UpdateRecurrence(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params recurrence.RecurrenceForm
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

		categoryID := &params.Category
		if *categoryID == 0 {
			categoryID = nil
		}

		_, err = db_sql.ChangeRecurrence(db, params.Id, params.Name, params.House_User, params.PayerAccountID, categoryID, params.Amount, params.Currency, params.ConversionRate, params.PayDate, params.DayCycle)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		recurrence, err := db_sql.GetRecurrence(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"recurrence": recurrence})
	}
}

//TODO: delete recurrence
