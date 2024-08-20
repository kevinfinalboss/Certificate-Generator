package main

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var tmpl *template.Template

func init() {
	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Key := os.Getenv("S3_TEMPLATE_KEY")

	obj, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		log.Fatalf("Failed to get template from S3: %v", err)
	}
	defer obj.Body.Close()

	htmlTemplate, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		log.Fatalf("Failed to read template body: %v", err)
	}

	tmpl, err = template.New("certificate").Parse(string(htmlTemplate))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := CertificateData{
		ParticipantName: request.QueryStringParameters["participant_name"],
		CourseName:      request.QueryStringParameters["course_name"],
		TotalHours:      request.QueryStringParameters["total_hours"],
		StartDate:       request.QueryStringParameters["start_date"],
		EndDate:         request.QueryStringParameters["end_date"],
		DirectorName:    request.QueryStringParameters["director_name"],
		CoordinatorName: request.QueryStringParameters["coordinator_name"],
		CertificateID:   request.QueryStringParameters["certificate_id"],
		IssueDate:       request.QueryStringParameters["issue_date"],
	}

	var filledTemplate bytes.Buffer
	if err := tmpl.Execute(&filledTemplate, data); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error",
		}, nil
	}

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
