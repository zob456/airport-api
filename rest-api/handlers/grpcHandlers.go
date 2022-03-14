package handlers

import (
	pb "AirportApi/airport-service/proto"
	"AirportApi/rest-api/models"
	"AirportApi/rest-api/utils"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func GetAirportDetails(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.AirportDetailsReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		grpcReq := &pb.AirportDetailsReq{AirportID: req.AirportID}

		standardCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		c := pb.NewAirportDataClient(conn)

		res, err := c.GetAirportDetails(standardCtx, grpcReq)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func GetAirportDistance(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *pb.AirportDistanceReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		standardCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		c := pb.NewAirportDataClient(conn)

		res, err := c.GetDistance(standardCtx, req)
		if err != nil {
			utils.AirHttpErrorHandler(ctx, err, http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}