# lowdb

## Usage

### Create a json file

```console
touch db.json
```

In file add:

```json
{}
```

### Save data

```go
package main

import (
	"github.com/valerianomacuri/lowdb/adapters"
	"github.com/valerianomacuri/lowdb/low"
)
// Define the data structure for database
type Data struct {
	Posts []string `json:"posts"`
}

func main() {
	adapter := adapters.NewJSONFile[Data]("db.json")
	db := low.New[Data](adapter)
    // Read database
	db.Read()
    // Write database in json
	defer db.Write()
	db.Data.Posts = append(db.Data.Posts, "hello world", "hola mundo", "olá mundo")
}
```
