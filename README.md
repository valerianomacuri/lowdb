# lowdb

> Simple to use local JSON database. Powered by Go. Inspired in lowdb for
> Javascript

```go
db.Data.Posts = append(
	db.Data.Posts, 
	Data.Posts{ ID: 1, Title: "lowdb is awesome"},
	)
```

```json
// db.json
{
  "posts": [
    { "id": 1, "title": "lowdb is awesome" }
  ]
}
```

## Install

```console
go get github.com/valerianomacuri/lowdb
```

## Usage

```console
touch db.json && echo "{}" > db.json
```

```json
// db.json
{}
```

```go
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
```

## Help us

[<img src="https://camo.githubusercontent.com/2e6f39f21c4d34ba1a4b01c34d27bb7e9dabb7d5c01fc8a4a2f1a571c36f4d5e/68747470733a2f2f63646e2e6275796d6561636f666665652e636f6d2f627574746f6e732f76322f64656661756c742d79656c6c6f772e706e673f773d33383426713d3735" width="200">](https://www.buymeacoffee.com/golowdb)
