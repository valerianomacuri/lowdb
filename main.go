package main

import (
	"github.com/valerianomacuri/lowdb/adapters"
	"github.com/valerianomacuri/lowdb/low"
)

type Data struct {
	Posts []string `json:"posts"`
}

func main() {
	adapter := adapters.NewJSONFile[Data]("db.json")
	db := low.New[Data](adapter)
	db.Read()
	defer db.Write()
	db.Data.Posts = append(db.Data.Posts, "hello world", "hola mundo", "olá mundo")
}
