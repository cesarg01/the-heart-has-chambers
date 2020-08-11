package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
type pageDetails struct {
	Title string
}

// Index page handler
func IndexPage(w http.ResponseWriter, r *http.Request) {

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

	titleHeader := pageDetails{Title: "the heart has chambers"}
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

	// Parse the html and check if there is an error.
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Add the struct slice and struct containing the title of the page into a interface mapped to m
	m := map[string]interface{}{
		"Data":      data,
		"PageTitle": titleHeader.Title,
	}

	// Execute template and pass in the struct and check if there is an error
	err = t.Execute(w, m)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
*/

// GetLogin credentials
func GetLogin(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// Updates the database from the xlsx
func UpdateDatabase() {
	// Open xlsx file to be read.
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the data from each row in the table.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		if row[0] == "Title" {
			continue
		}
		fmt.Println("Title", row[0])

		add(row)
		//fmt.Println(reflect.TypeOf(row))
		fmt.Println()
	}
}

// Render page. tmpl is the location of the file ex. templates/index.html
func Render(w http.ResponseWriter, tmpl string, data []Blog) {
	// Parse the html and check if there is an error.
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Execute template and pass in the struct and check if there is an error
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {

	// Used to get css on the pages.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	/*
		http.HandleFunc("/complete/", CompleteTaskFunc)
		http.HandleFunc("/delete/", DeleteTaskFunc)
		http.HandleFunc("/deleted/", ShowTrashTaskFunc)
		http.HandleFunc("/trash/", TrashTaskFunc)
		http.HandleFunc("/edit/", EditTaskFunc)
		http.HandleFunc("/completed/", ShowCompleteTasksFunc)
		http.HandleFunc("/restore/", RestoreTaskFunc)
		http.HandleFunc("/add/", AddTaskFunc)
		http.HandleFunc("/update/", UpdateTaskFunc)
		http.HandleFunc("/search/", SearchTaskFunc)
		http.HandleFunc("/login", GetLogin)
		http.HandleFunc("/register", PostRegister)
		http.HandleFunc("/admin", HandleAdmin)
		http.HandleFunc("/add_user", PostAddUser)
		http.HandleFunc("/change", PostChange)
		http.HandleFunc("/logout", HandleLogout)
	*/
	http.HandleFunc("/login/", GetLogin)
	http.HandleFunc("/", IndexPage)
	//http.HandleFunc("/work", WorkPage)
	http.HandleFunc("/about", AboutPage)

	//http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
