package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

	db, err := sql.Open("mysql", "root:thisiscool@/taskmanager_schema")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch true {
	case strings.HasPrefix(r.URL.Path, "/add"):
		AddTask(w, r, db)
		return
	case strings.HasPrefix(r.URL.Path, "/setStatus"):
		SetStatus(w, r, db)
		return
	case strings.HasPrefix(r.URL.Path, "/deleteTask"):
		DeleteTask(w, r, db)
		return
	}
}

// AddTask !
func AddTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/add")

	title := r.FormValue("title")
	now := time.Now()
	setAsFalse := 0

	q := fmt.Sprintf("INSERT INTO Task(Title, DateAdded, IsDone) VALUES('%s', '%v', '%v')", title, now.Format("2006/01/02T15:04:05"), setAsFalse)
	_, err := db.Query(q)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("New Task Added: '%s'", title))
}

// SetStatus !
func SetStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/setStatus")
	status := r.FormValue("status")
	taskID := r.FormValue("id")

	var q string
	if status != "" && taskID != "" {

		q = fmt.Sprintf("UPDATE Task SET IsDone = %s WHERE TaskID = %s", status, taskID)
		_, err := db.Query(q)
		if err != nil {
			log.Fatal(err)
		}

	}

}

// DeleteTask !
func DeleteTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/deleteTask")

	taskID := r.FormValue("id")
	q := fmt.Sprintf("DELETE FROM task WHERE TaskID = %v", taskID)
	_, err := db.Query(q)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted")
}
