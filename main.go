package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	// Connect to PostgreSQL database using gorm
	dsn := "host=localhost user=<please add USERNAME here> dbname=testdatabase port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate creates the messages table if it does not exist
	err = db.AutoMigrate(&Message{})
	if err != nil {
		panic(err)
	}
}

// Message is a simple struct to represent the data model
type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	Link    string `gorm:"type:text"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Golang backend!")
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMessages(w, r)
		case http.MethodPost:
			addMessage(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	if err := db.Model(&Message{}).Find(&messages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the array of messages to JSON and send it as the response
	responseJSON, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create shorturl
	randomText := GenerateRandomText(6)
	message.Link = fmt.Sprintf("%s%s", "https://shorturl/", randomText)

	// Add the new message to the database
	if err := db.Create(&message).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
