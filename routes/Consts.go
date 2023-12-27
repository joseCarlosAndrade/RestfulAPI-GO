package routes

import (
	"github.com/gorilla/mux"
)

var (
	Router *mux.Router
)

/* STRUCTS TAGS FOR JSON READING */
type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unity    string `json:"unity"`
}

type StoreProducts struct {
	Products []Product `json:"products"`
}
