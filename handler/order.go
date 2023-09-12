package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Create order"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("List order"))
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("GetById order"))
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Update order"))
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Delete order"))
}
