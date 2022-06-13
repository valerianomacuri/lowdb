package main

import (
	"github.com/valerianomacuri/lowdb/adapters"
	"github.com/valerianomacuri/lowdb/low"
)

// Define the data structure for database
type Data struct {
	Posts []Post `json:"posts"`
}
type Post struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

func main() {
	adapter := adapters.NewJSONFile[Data]("db.json")
	db := low.New[Data](adapter)
	// Read data from JSON file, this will set db.Data content
	db.Read()

	// Finally write db.Data content to file
	defer db.Write()

	// Append a post into posts array
	db.Data.Posts = append(
		db.Data.Posts,
		Post{ID: 1, Title: "lowdb is awesome"},
	)
}
