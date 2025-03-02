package main

import (
	"fmt"
	"lambda-func/app"

	"lambda-func/middleware"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Take some payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is required")
	}
	return fmt.Sprintf("Successfully processed user: %s", event.Username), nil
}

func ProtectedHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: "This is a protected route",
	}, nil
}

func main() {	
	myApp := app.NewApp()
	lambda.Start(func (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)  {
		switch request.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUser(request)
		case "/login":
			return myApp.ApiHandler.LoginUser(request)
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(request)
		default:
			return events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body: "Not Found",
			}, nil
		}
	})
}
