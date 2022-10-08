package main

import (
	"bin_schedule_v2/helpers"
	"context"
	_ "embed"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func lambdaGinRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Param("date")
		bins := csvSchedule[date]
		c.JSON(http.StatusOK, gin.H{
			"bins": bins,
		})
	}
}

func localGinRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Param("date")
		bins := csvSchedule[date]
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.JSON(http.StatusOK, gin.H{
			"bins": bins,
		})
	}
}

//go:embed test_schedule.csv
var schedule string

var csvSchedule, err = helpers.LoadFromCSV(schedule)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin started")
	server := gin.Default()

	server.GET("/bins/:date", lambdaGinRoute())
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginLambda = ginadapter.NewV2(server)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
