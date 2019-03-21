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
	port := "2223"
	log.Println("Listening port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}

func initRoutes() {
	http.HandleFunc("/auth", handlers.PerformLogin)                       //OK
	http.HandleFunc("/change_password", handlers.ResetPasswordToken)      //OK
	http.HandleFunc("/reset_password_email", handlers.ResetPasswordEmail) //OK
	http.HandleFunc("/create_task", handlers.CreateTask)                  //OK
	http.HandleFunc("/tasks_prof", handlers.GetTasksByProfessor)          //OK
	http.HandleFunc("/tasks_student", handlers.GetTasksForStudents)       //OK
	http.HandleFunc("/option_number", handlers.GetOptionNumberForTask)    //OK
	http.HandleFunc("/delete_task", handlers.DeleteTask)                  //OK
	http.HandleFunc("/forums", handlers.GetForums)                        //OK
	http.HandleFunc("/forum", handlers.GetForum)                          //OK
	http.HandleFunc("/send_message", handlers.SendMessage)                //OK
	http.HandleFunc("/delete_message", handlers.DeleteMessage)            //OK
	http.HandleFunc("/registration", handlers.Registration)               //OK
	http.HandleFunc("/check_session", handlers.CheckSession)              //OK
	http.HandleFunc("/groups", handlers.GetGroups)                        //OK
	http.HandleFunc("/all_groups", handlers.GetAllGroups)                 //OK
	http.HandleFunc("/add_group", handlers.AddGroup)                      //OK
	http.HandleFunc("/delete_group", handlers.DeleteGroup)                //OK
	http.HandleFunc("/download_prof", handlers.DownloadTaskForProf)       //OK
	http.HandleFunc("/download_student", handlers.DownloadTaskForStudent) //OK
}
