package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

// IndexPage ...
func AboutPage(w http.ResponseWriter, r *http.Request) {

	// Load database url using environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//fmt.Println(os.Getenv("USER"))

	// Connecting to database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Printf("Connection worked!\n")

	rows, _ := conn.Query(context.Background(), "SELECT * FROM blog ORDER BY created_date DESC LIMIT 5;")

	// To use Blog struct must run the blog_model.go file with main.go
	// Create a struct slice
	var data []Blog
	for rows.Next() {
		var blg Blog
		err := rows.Scan(&blg.ID, &blg.Author, &blg.Title, &blg.Slug, &blg.Body, &blg.CreatedDate, &blg.Published)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(blg.ID, blg.Author, blg.Title, blg.Slug, blg.Body, blg.CreatedDate.Format("2006-01-02"), blg.Published)
		data = append(data, blg)
	}
	// Print out the slice full of structs
	fmt.Println(data)

	Render(w, "templates/about.html", data)
}
