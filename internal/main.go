package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jiminhu/LeetCode/LeetCode/systemdesign/restapis/pkg/swagger/server/restapi"
	"github.com/jiminhu/LeetCode/LeetCode/systemdesign/restapis/pkg/swagger/server/restapi/operations"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSignupAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// Implement the SignUp handler
	api.CreateUserHandler = operations.CreateUserHandlerFunc(
		func(user operations.CreateUserParams) middleware.Responder {
			return operations.CreateUserOK().WithPayload("OK" + "user id")
		})

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
