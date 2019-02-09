package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

// Database instance
var db *sql.DB
var key []byte

// Initialize
func init() {
	file, err := ioutil.ReadFile("./properties.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	type dbCredentials struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		DBName   string `json:"db_name"`
		User     string `json:"user"`
		Password string `json:"password"`
		Key      string `json:"key"`
	}
	var credentials dbCredentials
	json.Unmarshal(file, &credentials)

	connection := fmt.Sprintf("host=%s port=%d dbname=%s user=%s "+
		"password=%s sslmode=disable",
		credentials.Host, credentials.Port, credentials.DBName, credentials.User, credentials.Password)

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	key = []byte(credentials.Key)
	return
}

func GenerateNewPassword() (password string, err error) {
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(25))
		bytes[i] = byte(65 + nBig.Int64())
	}
	password = string(bytes)
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
