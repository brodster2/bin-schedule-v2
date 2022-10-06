package main

import (
	"context"
	"encoding/csv"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
	"strings"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func LoadFromCSV(filepath string) (map[string][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil { // Discard column header
		return nil, err
	}

	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	schedule := make(map[string][]string)

	for _, v := range content {
		var strip []string
		for _, s := range strings.Split(v[1], ":") {
			strip = append(strip, strings.TrimSpace(s))
		}
		schedule[v[0]] = strip
	}

	return schedule, nil
}

func lambdaGinRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Param("date")
		bins := schedule[date]
		c.JSON(http.StatusOK, gin.H{
			"bins": bins,
		})
	}
}

func localGinRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Param("date")
		bins := schedule[date]
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.JSON(http.StatusOK, gin.H{
			"bins": bins,
		})
	}
}

var schedule, err = LoadFromCSV("test_schedule.csv")

var ginLambda *ginadapter.GinLambda
var server *gin.Engine

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin started")
	server = gin.Default()

	var routeHandler func(c *gin.Context)
	exeContext := os.Getenv("APP_CONTEXT")

	if exeContext == "lambda" {
		routeHandler = lambdaGinRoute()
	} else {
		routeHandler = localGinRoute()
	}

	server.GET("/bins/:date", routeHandler)

	ginLambda = ginadapter.New(server)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	exeContext := os.Getenv("APP_CONTEXT")

	if exeContext == "lambda" {
		lambda.Start(Handler)
	} else {
		server.Run() // listen and serve on 0.0.0.0:8080
	}
}
