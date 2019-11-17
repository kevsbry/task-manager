package main

import (
	"fmt"
	"net/http"

	"github.com/kevsbry/taskManager/api"
	"github.com/kevsbry/taskManager/handler"
)

func main() {
	//file handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	const PORT = ":9999"
	http.HandleFunc("/", handler.TaskManager)
	http.HandleFunc("/api/", api.Handler)

	go fmt.Println(fmt.Sprintf("server running on port%s", PORT))
	http.ListenAndServe(PORT, nil)
}
