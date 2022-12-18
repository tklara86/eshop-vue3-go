package main

import (
	"context"
	"log"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tklara86/eshop-vue3-go/internal/server"
)

var ginLambda *ginadapter.GinLambda

func init() {
	s, err := server.New(server.Config{
		Port: 9090,
	})
	if err != nil {
		log.Fatalf("server not created. Exited with error: %s", err)
	}
	ginLambda = ginadapter.New(s.Engine)

}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, event)
}

func main() {
	lambda.Start(Handler)
}
