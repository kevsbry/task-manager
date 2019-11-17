package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/dustin/go-humanize"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kevsbry/taskManager/models"
)

// TaskManager !
func TaskManager(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:thisiscool@/taskmanager_schema?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch true {
	case strings.HasPrefix(r.URL.Path, "/task"):
		TaskPage(w, r, db)
		return
	}

}

// TaskPage !
func TaskPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/task")

	task := []models.Task{}
	tempTask := models.Task{}
	type Context struct {
		TaskID        int
		TaskTitle     string
		TaskDateAdded string
		TaskIsDone    bool
	}
	context := []Context{}

	result, err := db.Query("SELECT * FROM Task")
	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		result.Scan(&tempTask.TaskID, &tempTask.Title, &tempTask.DateAdded, &tempTask.IsDone)
		task = append(task, tempTask)
	}

	for _, t := range task {
		context = append(context, Context{
			TaskID:        t.TaskID,
			TaskTitle:     t.Title,
			TaskDateAdded: humanize.Time(t.DateAdded), // had to create context because of humanizing time
			TaskIsDone:    t.IsDone,
		})
	}

	t := template.New("")
	t, _ = t.ParseFiles("static/ui/html/task.html")
	_ = t.ExecuteTemplate(w, "task.html", context)
}
