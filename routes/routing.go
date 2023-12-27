/*
# Routing module

Módulo que impĺementa as funções handlers de cada endpoint da API.
*/
package routes

import (
	// "fmt"
	// "github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	// "encoding/json"
)

/* Handler function to / route (root)*/
func RootRoute(w http.ResponseWriter, r * http.Request) {
	http.ServeFile(w, r, "static/index.html")
}


/* GET Methods */
func GetProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get method requested")
	// fmt.Fprintf(w, "Retornando get!!!!!!")
	// vars := mux.Vars(r)
	// fmt.Println(vars)

	q := r.URL.Query()
	name := q.Get("name")

	if name != "" {
		// fmt.Fprint(w, name)
		fmt.Println("Getting a request for product : ", name)

		storeProductsJSON, err := os.Open("data/store.json")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"status" : "Internal server error! File unavailable"}`)
			return
		}

		defer storeProductsJSON.Close()

		byteValue, err := io.ReadAll(storeProductsJSON)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"status" : "Internal server error! %v"}`, err.Error())
			return
		}

		var products StoreProducts
		json.Unmarshal(byteValue, &products)

		for _, v := range products.Products {
			if name == v.Name {
				
				response := fmt.Sprintf(`{"status" : "OK! Product found" ,` + 
					`"product" : {`+
						`"name" : "%s",`+
						`"quantity" : "%d",`+
						`"unity" : "%s"`+
					`}`+
				`}`, v.Name, v.Quantity, v.Unity)
				fmt.Fprint(w, response)
				return
			}
		}
		fmt.Fprintf(w, `{"status" : "OK! Product not found"}`)

	} else {
		fmt.Println("sending bad request")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "status" : "Bad request! Invalid variables"}`)
	}
}

/* POST Methods */


//////////////////////////// planejamento //////////////

/*
	CONSULTAR PRODUTO [GET] /get_product?name=nomedoproduto
		RESPOSTA:
		400:
			404 - produto nao encontrado (inexistente)

		200:
			{ 
				"product" : {
					"name" : "...",
					"quantity" : 10,
					"price" : 14.11	
				}
			} 
*/