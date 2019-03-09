package main

import (
	"MSA/handlers"
	"MSA/testing"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

	// ДЛЯ ЗАГЛУШЕК ОТ БД
	testing.SetTestMode(false)

	http.Handle("/", http.FileServer(http.Dir("assets")))
	initRoutes()
	port := "8080"
	log.Println("Listening port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}

func initRoutes() {
	http.HandleFunc("/auth", handlers.PerformLogin)
	http.HandleFunc("/reset_password", handlers.ResetPassword)
	http.HandleFunc("/create_task", handlers.CreateTask)
	http.HandleFunc("/forums", handlers.GetForums)
	http.HandleFunc("/forum", handlers.GetForum)
	http.HandleFunc("/send_message", handlers.SendMessage)
	http.HandleFunc("/registration", handlers.Registration)
	http.HandleFunc("/check_session", handlers.CheckSession)
}
