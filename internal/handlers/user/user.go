package userHandlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Rombond/budgestify/api/types/user"
	password "github.com/Rombond/budgestify/internal/password"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"

	"github.com/gin-gonic/gin"
)

func GetUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		valid, err := token.IsTokenValid(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Authorization is not valid"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		user, err := db_sql.GetUser(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func CreateUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params user.UserForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		hash, err := password.ParamToByte(params.Hash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		id, err := db_sql.CreateUser(db, params.Name, params.Login, hash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"userID": id})
	}
}

func LoginUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params user.UserForm
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.Login == "" && params.Hash == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "Login or Password is empty"})
			return
		}

		hash, err := password.ParamToByte(params.Hash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		id, err := db_sql.GetUserID(db, params.Login)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		user, err := db_sql.GetUser(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		valid, err := password.IsPasswordValid(user.Hash, hash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Wrong password"})
			return
		}

		tokenStr, err := token.GenerateToken(params.Login)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenStr})
	}
}

func ChangeUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		valid, err := token.IsTokenValid(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"reason": "Authorization is not valid"})
			return
		}

		var params user.UserForm
		err = ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.Login == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No Login provided"})
			return
		}
		id := params.Id
		if id == 0 {
			id, err = db_sql.GetUserID(db, params.Login)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		if params.Hash != "" {
			hash, err := password.ParamToByte(params.Hash)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
			_, err = db_sql.ChangeUserHash(db, id, hash)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		if params.Name != "" {
			_, err = db_sql.ChangeUserName(db, id, params.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		user, err := db_sql.GetUser(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
