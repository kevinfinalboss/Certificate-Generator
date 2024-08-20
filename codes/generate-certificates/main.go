package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

var (
	tmpl        *template.Template
	s3Svc       *s3.S3
	s3Bucket    string
	dynamoDBSvc *dynamodb.DynamoDB
	tableName   string
)

func init() {
	sess := session.Must(session.NewSession())
	dynamoDBSvc = dynamodb.New(sess)
	s3Svc = s3.New(sess)

	tableName = os.Getenv("DYNAMODB_TABLE_NAME")
	s3Bucket = os.Getenv("S3_BUCKET_NAME")
	s3Key := os.Getenv("S3_TEMPLATE_KEY")

	log.Printf("Loading HTML template from S3 bucket: %s, key: %s", s3Bucket, s3Key)
	obj, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		log.Fatalf("Failed to get template from S3: %v", err)
	}
	defer obj.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(obj.Body)
	if err != nil {
		log.Fatalf("Failed to read template content: %v", err)
	}
	htmlTemplate := buf.String()

	tmpl, err = template.New("certificate").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
	log.Println("HTML template loaded and parsed successfully")
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %v", request)

	switch request.HTTPMethod {
	case "POST":
		return handlePost(ctx, request)
	case "GET":
		return handleGet(ctx, request)
	default:
		log.Printf("Invalid HTTP method: %s", request.HTTPMethod)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Method Not Allowed",
		}, nil
	}
}

func handlePost(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Handling POST request")

	var data CertificateData
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid request body",
		}, nil
	}

	if data.UUID == "" {
		data.UUID = uuid.New().String()
	}

	log.Printf("Generated UUID: %s", data.UUID)

	item, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		log.Printf("Failed to marshal item: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to marshal item",
		}, nil
	}

	if _, ok := item["UUID"]; !ok {
		item["UUID"] = &dynamodb.AttributeValue{S: aws.String(data.UUID)}
		log.Printf("Manually added UUID to the item: %s", data.UUID)
	} else {
		log.Printf("UUID after marshalling: %v", item["UUID"])
	}

	_, err = dynamoDBSvc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		log.Printf("Failed to save item in DynamoDB: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to save item in DynamoDB",
		}, nil
	}

	log.Printf("Certificate created successfully with UUID: %s", data.UUID)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Certificate created with UUID: %s", data.UUID),
	}, nil
}

func handleGet(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Handling GET request")

	uuid := request.QueryStringParameters["uuid"]
	if uuid == "" {
		log.Println("UUID parameter is missing")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "UUID parameter is required",
		}, nil
	}

	log.Printf("Fetching certificate with UUID: %s", uuid)
	result, err := dynamoDBSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(uuid),
			},
		},
	})
	if err != nil {
		log.Printf("Failed to retrieve item from DynamoDB: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to retrieve item from DynamoDB",
		}, nil
	}

	if result.Item == nil {
		log.Printf("No item found with UUID: %s", uuid)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Certificate not found",
		}, nil
	}

	var data CertificateData
	err = dynamodbattribute.UnmarshalMap(result.Item, &data)
	if err != nil {
		log.Printf("Failed to unmarshal item: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to unmarshal item",
		}, nil
	}

	var filledTemplate bytes.Buffer
	if err := tmpl.Execute(&filledTemplate, data); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error",
		}, nil
	}

	log.Printf("Certificate successfully retrieved for UUID: %s", uuid)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       filledTemplate.String(),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
