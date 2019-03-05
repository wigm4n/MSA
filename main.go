package main

import (
	"MSA/handlers"
	"MSA/sampling"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	initRoutes()
	log.Println("Listening port 8080...")
	http.ListenAndServe(":8080", nil)

	//fmt.Println(sampling.ReturnTask1(10, 1, 100, 10, 0.05))
	//fmt.Println(sampling.ReturnTask2(10, 1, 100, 10, 0.05))
	//fmt.Println(sampling.ReturnTask3(10, 1, 100, 10, 0.05))
	//fmt.Println(sampling.ReturnTask4(10,  0.05))
	//fmt.Println(sampling.ReturnTask5(10, 1, 100, 10, 0.05))
	fmt.Println(sampling.ReturnTask6(10, 1, 100, 10, 10, 10, 0.05))

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
