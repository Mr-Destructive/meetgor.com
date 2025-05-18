package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mr-destructive/mr-destructive.github.io/plugins"
	"github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		ctx := context.Background()
		dbName := os.Getenv("TURSO_DATABASE_NAME")
		dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

		var err error
		dbString := fmt.Sprintf("libsql://%s?authToken=%s", dbName, dbToken)
		db, err := sql.Open("libsql", dbString)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
		}
		defer db.Close()

		queries := libsqlssg.New(db)
		if _, err := db.ExecContext(ctx, plugins.DDL); err != nil {
			return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
		}

		var payload libsqlssg.CreateAuthorParams
		err = json.Unmarshal([]byte(req.Body), &payload)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Invalid Payload"), nil
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Invalid Payload"), nil
		}
		payload.Password = string(hashedPassword)
		userID, err := queries.CreateAuthor(ctx, payload)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "User creation failed"), nil
		}
		user, err := queries.GetAuthorByID(ctx, userID)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "User Not Found"), nil
		}
		return jsonResponse(http.StatusOK, map[string]string{
			"username": user.Username,
			"id":       string(user.ID),
		}), nil
	default:
		return errorResponse(http.StatusMethodNotAllowed, "Method Not Allowed"), nil
	}

}

func Authenticate(username, rawPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(rawPassword), []byte(hashedPassword))
	fmt.Println(err)
	if err != nil {
		fmt.Println("Authentication Failure")
		return false
	}
	return true
}
func jsonResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func errorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return jsonResponse(statusCode, map[string]string{"error": message})
}
