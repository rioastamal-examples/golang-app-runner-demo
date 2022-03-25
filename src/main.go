// A very simple demo of Go based web app
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"net/http"
	"os"
)

type Item struct {
	Pk   string `json:"pk"`
	Sk   string `json:"sk"`
	Data string `json:"data"`
}

func checkAuth(response http.ResponseWriter, request *http.Request) (map[string]string, error) {
	credentials := map[string]string{
		"username": "",
		"password": "",
	}
	user, passwd, ok := request.BasicAuth()

	if !ok {
		response.Header().Set("WWW-Authenticate", "Basic realm=\"Go Basic Auth\"")
		response.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintf(response, "Access denied!")
		return credentials, errors.New("Access denied!")
	}

	usernameAndPassword := user + passwd
	if usernameAndPassword != os.Getenv("APP_USERNAME")+os.Getenv("APP_PASSWORD") {
		response.Header().Set("WWW-Authenticate", "Basic realm=\"Go Basic Auth\"")
		response.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintf(response, "Invalid username or password!")
		return credentials, errors.New("Invalid username or password")
	}

	credentials["username"] = user
	credentials["password"] = passwd

	return credentials, nil
}

func saveProvidersConfig(client *dynamodb.Client, item map[string]string) {
	userItem := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"pk":   &types.AttributeValueMemberS{Value: item["pk"]},
			"sk":   &types.AttributeValueMemberS{Value: item["sk"]},
			"data": &types.AttributeValueMemberS{Value: item["data"]},
		},
		TableName: aws.String(os.Getenv("APP_TABLE_NAME")),
	}
	client.PutItem(context.TODO(), userItem)
}

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	ddbClient := dynamodb.NewFromConfig(cfg)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		credentials, authErr := checkAuth(response, request)
		if authErr != nil {
			return
		}

		if request.Method == "POST" && request.URL.Path == "/" {
			formData := map[string]string{
				"pk":   fmt.Sprintf("user#%s", credentials["username"]),
				"sk":   fmt.Sprintf("data#%s", credentials["username"]),
				"data": request.FormValue("providers-config"),
			}

			saveProvidersConfig(ddbClient, formData)
		}

		http.FileServer(http.Dir("./static")).ServeHTTP(response, request)
	})

	http.HandleFunc("/api/v1/providers/", func(response http.ResponseWriter, request *http.Request) {
		credentials, authErr := checkAuth(response, request)
		if authErr != nil {
			return
		}

		paramItem := &dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: fmt.Sprintf("user#%s", credentials["username"])},
				"sk": &types.AttributeValueMemberS{Value: fmt.Sprintf("data#%s", credentials["username"])},
			},
			TableName: aws.String(os.Getenv("APP_TABLE_NAME")),
		}
		resp, _ := ddbClient.GetItem(context.TODO(), paramItem)

		itemData := Item{}
		attributevalue.UnmarshalMap(resp.Item, &itemData)

		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)

		fmt.Fprintf(response, "%s", itemData.Data)
	})

	http.HandleFunc("/api/v1/whoami/", func(response http.ResponseWriter, request *http.Request) {
		credentials, authErr := checkAuth(response, request)
		if authErr != nil {
			return
		}

		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, `{ "username": "%s", "auth_data": "%s"}`, credentials["username"], request.Header.Get("Authorization"))
	})

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Print("Running web server on :", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
