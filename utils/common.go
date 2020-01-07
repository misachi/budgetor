package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "budgetor/users"
)

// Conn handle to connection to database
var Conn = initializeDb()

func initializeDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return db
}

type Entity interface {
	DBMigrate(db *gorm.DB)
}

type ErrorResponse struct {
	Error string
}

func JsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response, _ := JSONMarshal(data)
	w.Write(response)
}

// JSONMarshal serialize to Json
func JSONMarshal(data interface{}) ([]byte, error) {
	var b []byte
	var err error

	b, err = json.MarshalIndent(data, "", "   ")

	return b, err
}
