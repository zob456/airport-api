package middleware

import (
	"AirportApi/rest-api/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

type JwtClaims struct {
	UserID    string
	UserName    string
	jwt.StandardClaims
}

var (
	secret = os.Getenv("SECRET")
)

func ValidateToken(db *sql.DB, airCtx *utils.AirContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL.Path)
		if ctx.Request.URL.Path == "/auth/login" {
			return
		}
		authToken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))
		token := strings.TrimPrefix(authToken, "Bearer ")
		if token == "null" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "not signed in")
			return
		}
		referer := ctx.Request.Header.Get("Referer")
		valid, claims := VerifyToken(token, referer)
		if !valid {
			ctx.JSON(http.StatusUnauthorized, fmt.Errorf("not authorized"))
			return
		}
		err := SelectAuthedUserLoggedInStatus(db, token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "not signed in")
			return
		}
		airCtx.UserID = claims.UserID
		airCtx.UserName = claims.UserName
		airCtx.JWT = token
	}
}

func VerifyToken(tokenString string, origin string) (bool, *JwtClaims) {
	if tokenString == "" {
		return false, nil
	}
	claims := &JwtClaims{}
	token, err := getTokenFromString(tokenString, claims)
	if err != nil {
		return false, nil
	}
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func getTokenFromString(tokenString string, claims *JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func SelectAuthedUserLoggedInStatus(db *sql.DB, jwt string) error {
	/*language=PostgreSQL*/
	const query = `SELECT "UserID"
		FROM "UserData"."user"
		WHERE "JWT" = $1 AND "Active" = true
		LIMIT 1;`

	var userID *string
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(jwt)
	err = row.Scan(
		&userID,
	)
	if err != nil {
		return err
	}
	return nil
}
