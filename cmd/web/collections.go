package web

import (
	"net/http"
)

type Task struct{}

func (o *Task) Create(w http.ResponseWriter, r *http.Request) {
	panic("Create: TODO")
}

func (o *Task) List(w http.ResponseWriter, r *http.Request) {
	panic("List: TODO")
}

func (o *Task) GetById(w http.ResponseWriter, r *http.Request) {
	panic("GetById: TODO")
}

func (o *Task) UpdateById(w http.ResponseWriter, r *http.Request) {
	panic("UpdateById: TODO")
}

func (o *Task) DeleteById(w http.ResponseWriter, r *http.Request) {
	panic("DeleteById: TODO")
}
