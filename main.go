package main

import (
	"MSA/handlers"
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

	http.Handle("/", http.FileServer(http.Dir("assets")))
	initRoutes()
	port := "8888"
	log.Println("Listening port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}

func initRoutes() {
	http.HandleFunc("/auth", handlers.PerformLogin)
	http.HandleFunc("/reset_password", handlers.ResetPassword)
	http.HandleFunc("/create_task", handlers.CreateTask)
	http.HandleFunc("/tasks_prof", handlers.GetTasksByProfessor)
	http.HandleFunc("/tasks_student", handlers.GetTasksForStudents)
	http.HandleFunc("/forums", handlers.GetForums)
	http.HandleFunc("/forum", handlers.GetForum)
	http.HandleFunc("/send_message", handlers.SendMessage)
	http.HandleFunc("/registration", handlers.Registration)
	http.HandleFunc("/check_session", handlers.CheckSession)
	//getTasks
	//getTask
	//deleteTask
	//getGroups
	//addGroup
	//download_file
	//изменить createtask на получение листов заданий

	//КАК ВЫДАВАТЬ СПИСОК ФОРУМОВ ДЛЯ СТУДЕНТА?
}
