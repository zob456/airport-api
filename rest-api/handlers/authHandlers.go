package handlers

import (
	"AirportApi/rest-api/data"
	"AirportApi/rest-api/middleware"
	"AirportApi/rest-api/models"
	"AirportApi/rest-api/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.LoginReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}
		authedUser, err := data.SelectAuthedUser(h.DB, req.UserName, req.Password)
		if err != nil {
			utils.AirSqlErrorHandler(ctx, err)
			return
		}

		if authedUser != nil {
			token, err := middleware.GenerateToken(&middleware.JwtClaims{
				UserID:         authedUser.UserID,
				UserName: authedUser.UserName,
				StandardClaims: jwt.StandardClaims{},
			}, time.Now().Add(time.Hour*720).UTC(), authedUser.UserID)
			if err != nil {
				log.Printf("Error creating JWT: %+v", err)
			}
			err = data.AddJWT(h.DB, token, authedUser.UserID)
			if err != nil {
				log.Printf("ERROR: %+v\n", err)
				ctx.JSON(http.StatusInternalServerError, "Failed to update token")
				return
			}
			authedUser.JWT = token

		} else {
			ctx.JSON(http.StatusUnauthorized, "Failed to authenticate")
			return
		}
		ctx.JSON(http.StatusOK, authedUser)
	}
}

func (h *AuthHandler) Logout(airCtx *utils.AirContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := airCtx.UserID
		err := data.DeleteJWT(h.DB, userID)
		if err != nil {
			log.Printf("ERROR: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, "Failed to remove sik")
			return
		}
		token := airCtx.JWT
		err = middleware.SelectAuthedUserLoggedInStatus(h.DB, token)
		if err != nil {
			utils.AirExpectedNoRowsInSqlErrorHandler(ctx, err)
		}

		ctx.JSON(http.StatusOK, "successfully signed out!")
	}
}
