package main

import (
	"copyPasteProj/internal/taskstore"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"time"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks/" {
		if r.Method == http.MethodPost {
			ts.createTaskHandler(w, r)
		} else if r.Method == http.MethodGet {
			ts.getAllTasksHandler(w, r)
		} else if r.Method == http.MethodDelete {
			ts.deleteAllTasksHandler(w, r)
		} else {
			http.Error(w, fmt.Sprintf("expected method get / post / delete, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}

	}
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling task create at %s\n", r.URL.Path)

	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	js, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (ts *taskServer) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get all tasks at %s\n", r.URL.Path)
	allTasks := ts.store.GetAllTasks()
	js, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func (ts *taskServer) deleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	mux.HandleFunc("/", server.taskHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}
