package main

import (
	// "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/joseCarlosAndrade/RestfulAPI-GO/routes"
	"log"
)

func main() {
	
	routes.Router = mux.NewRouter()

	routes.Router.HandleFunc("/get_product", routes.GetProduto).Methods("GET")
	routes.Router.HandleFunc("/", routes.RootRoute)
	
	// serving page
	log.Println("Listening on localhost:3000 ")
	http.ListenAndServe(":3000", routes.Router)
}

/* example
func main() {
	// Create a new router
	router := mux.NewRouter()

	// Serve static files from the "static" directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define a handler function for API GET requests
	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Hello, this is the API!"}`)
	}

	// Route for API requests
	router.HandleFunc("/api", apiHandler).Methods("GET")

	// Route for serving the static HTML page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
*/


/*
anotações importantes :
variaveis de url na chamada get

	routes.Router.HandleFunc("/get_product/{name}", routes.GetProduto).Methods("GET")

^^ faz com que o request localhost:3000/get_product/variavel_aqui seja possivel contendo qualquer coisa em variavel_aqui

caso seja desejavel url query (?name=nome_aqui):
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
*/
