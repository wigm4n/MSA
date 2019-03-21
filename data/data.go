package data

import (
	zipimpl "archive/zip"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"
	"path/filepath"
)

// Database instance
var db *sql.DB
var key []byte

// Initialize
func init() {
	file, err := ioutil.ReadFile("./properties_db.json")
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

func GenerateNewPassword() (password string) {
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

func GenerateSessionToken() (password string) {
	bytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(25))
		bytes[i] = byte(65 + nBig.Int64())
	}
	password = string(bytes)
	return
}

func MakeArchive(pathToTaskDirectory string) (error, string, string) {
	inFilePath := pathToTaskDirectory[2:] + "/Data"
	outFilePath := pathToTaskDirectory + "/ProfData.zip"
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err, "", ""
	}
	defer func() {
		_ = outFile.Close()
	}()

	return Archive(inFilePath, outFile), outFilePath, inFilePath + "/Tasks"
}

func Archive(inFilePath string, writer io.Writer) error {
	zipWriter := zipimpl.NewWriter(writer)

	basePath := filepath.Dir(inFilePath)

	err := filepath.Walk(inFilePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil || fileInfo.IsDir() {
			return err
		}

		relativeFilePath, err := filepath.Rel(basePath, filePath)
		if err != nil {
			return err
		}

		archivePath := path.Join(filepath.SplitList(relativeFilePath)...)

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer func() {
			_ = file.Close()
		}()

		zipFileWriter, err := zipWriter.Create(archivePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFileWriter, file)
		return err
	})
	if err != nil {
		return err
	}

	return zipWriter.Close()
}

type ProgressFunc func(archivePath string)
