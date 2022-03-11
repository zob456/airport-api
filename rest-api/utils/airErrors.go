package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

const PublicBadRequest string = "malformed payload"
const PublicNoRowsInSQL string = "requested data not found"
const PublicInternalServerError string = "Internal Server Error"
const PublicNotAuthorized string = "not authorized"

func ReturnPublicErrorMessage(errCode int) string {
	var err string
	switch errCode {
	case 400:
		err = PublicBadRequest
	case 401:
		err = PublicNotAuthorized
	default:
		err = PublicInternalServerError
	}
	return err
}

func AirHttpErrorHandler(ctx *gin.Context, err error, errorCode int) {
	publicErrMessage := ReturnPublicErrorMessage(errorCode)
	log.Println(fmt.Sprintf("ERROR: %+v", err))
	debug.PrintStack()
	ctx.AbortWithStatusJSON(errorCode, publicErrMessage)
	return
}

func AirSqlErrorHandler(ctx *gin.Context, err error) {
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println(fmt.Sprintf("ERROR: %+v", err))
			ctx.AbortWithStatusJSON(404, PublicNoRowsInSQL)
			return
		} else {
			log.Println(fmt.Sprintf("ERROR: %+v", err))
			ctx.AbortWithStatusJSON(500, PublicInternalServerError)
			return
		}
	}
}

func AirExpectedNoRowsInSqlErrorHandler(ctx *gin.Context, err error) {
	if err.Error() != "sql: no rows in result set" {
		log.Println(fmt.Sprintf("ERROR: %+v", err))
		ctx.AbortWithStatusJSON(500, PublicInternalServerError)
		return
	}
}
