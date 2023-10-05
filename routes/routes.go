package routes

import (
	"log"
	"net/http"

	"github.com/adrielldev/api-rest-go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
