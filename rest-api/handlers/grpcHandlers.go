package handlers

import (
	pb "AirportApi/airport-service/proto"
	"AirportApi/rest-api/models"
	"AirportApi/rest-api/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAirportDetails(client pb.AirportDataClient, standardCtx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.AirportDetailsReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		grpcReq := &pb.AirportDetailsReq{AirportID: req.AirportID}

		res, err := client.GetAirportDetails(standardCtx, grpcReq)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func GetAirportDistance(client pb.AirportDataClient, standardCtx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.AirportDistanceReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		grpcReq := &pb.AirportDistanceReq{FirstAirportID: req.FirstAirportID, SecondAirportID: req.SecondAirportID}

		res, err := client.GetDistance(standardCtx, grpcReq)

		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}