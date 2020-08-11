package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

func add(dataInfo []string) {
	// Load database url using environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecting to database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Printf("Connection worked!\n")
	fmt.Println(dataInfo[3])
	sqlStatement := `INSERT INTO blog (author, title, slug, body, created_date, published) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = conn.Exec(context.Background(), sqlStatement, "Olivia R.", dataInfo[0], dataInfo[1], dataInfo[2], dataInfo[3], "false")

	if err != nil {
		panic(err)
	}
}
