// main.go
package main

import (
	"database/sql"
	_ "fmt"
	_ "io"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
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

	err = sqltocsv.Write(os.Stdout, rows)
	if err != nil {
		panic(err)
	}

	return `{"response":"Hello ƛ!"}`, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
