package houseHandlers

import (
	"database/sql"
	"net/http"

	"github.com/Rombond/budgestify/api/types/house"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/Rombond/budgestify/internal/token"
	"github.com/gin-gonic/gin"
)

func GetHouse(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params house.HouseForm
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

		house, err := db_sql.GetHouse(db, params.HouseID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		house.UserId = params.UserId
		ctx.JSON(http.StatusOK, house)
	}
}

func CreateHouseForUser(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params house.HouseForm
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

		created := false
		if params.Name != "" && params.HouseID == 0 {
			params.HouseID, err = db_sql.CreateHouse(db, params.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
			created = true
		} else if params.Name == "" && params.HouseID == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "Need 'name' to create or 'houseID' to link an house to this user"})
			return
		} else {
			_, err := db_sql.GetHouse(db, params.HouseID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		if created {
			_, err = db_sql.CreateHouseUser(db, params.HouseID, params.UserId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
			ctx.JSON(http.StatusCreated, gin.H{"status": "New house created and linked to user as administrator"})
		} else {
			isInside, err := db_sql.IsUserInThisHouse(db, params.HouseID, params.UserId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
			if isInside {
				ctx.JSON(http.StatusConflict, gin.H{"reason": "User is already a member of this house"})
			} else {
				_, err = db_sql.InviteUserToHouse(db, params.HouseID, params.UserId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"status": "House founded and linked to user as member"})
			}
		}
	}
}

func ChangeHouse(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var params house.House
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if params.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No userID provided"})
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

		if params.Id == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": "No id provided"})
			return
		}

		if params.Name != "" {
			_, err = db_sql.ChangeHouseName(db, params.Id, params.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
				return
			}
		}

		house, err := db_sql.GetHouse(db, params.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, house)
	}
}

//TODO: delete house
