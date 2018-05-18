// main.go
package main

import (
	"database/sql"
	"fmt"
	_ "io"
	"log"
	"os"
	"bytes"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
)

func hello() (string, error) {

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM public.django_migrations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var writer bytes.Buffer
	err = sqltocsv.Write(&writer, rows)
	//	os.Stdout, rows)


	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	uploader := s3manager.NewUploader(sess)
	//f, err  := os.Open(filename)

// Upload the file to S3.
result, err := uploader.Upload(&s3manager.UploadInput{
    Bucket: aws.String("dness-temp"),
    Key:    aws.String("csv_from_fable.csv"),
    Body:   &writer,
})
if err != nil {
    panic(err)
}
fmt.Println(result)


	
	if err != nil {
		panic(err)
	}

	return `{"response":"Hello ƛ!"}`, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
