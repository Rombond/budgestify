package categoryHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/category"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetCategory(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params category.CategoryForm
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

		category, err := db_sql.GetCategory(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, category)
	}
}

func GetCategories(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params category.CategoryForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		if params.HouseId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No houseID provided"})
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

		categories, err := db_sql.GetCategories(db, params.HouseId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, categories)
	}
}

func CreateCategory(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params category.CategoryForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
			return
		}

		if params.HouseId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No houseID provided"})
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

		id, err := db_sql.CreateCategory(db, params.Name, params.Icon, params.Parent, params.HouseId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"categoryID": id})
	}
}

func UpdateCategory(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params category.CategoryForm
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

		_, err = db_sql.ChangeCategory(db, params.Id, params.Name, params.Icon, params.Parent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		category, err := db_sql.GetCategory(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"category": category})
	}
}

//TODO: delete category
