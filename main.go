package main

import (
	"MSA/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	initRoutes()
	log.Println("Listening port 8080...")
	http.ListenAndServe(":8080", nil)
}

func initRoutes() {
	http.HandleFunc("/auth", handlers.PerformLogin)
	http.HandleFunc("/reset_password", handlers.ResetPassword)
	http.HandleFunc("/create_task", handlers.CreateTask)
	http.HandleFunc("/forums", handlers.GetForums)
	http.HandleFunc("/forum", handlers.GetForum)
	http.HandleFunc("/send_message", handlers.SendMessage)
	http.HandleFunc("/registration", handlers.Registration)
}
