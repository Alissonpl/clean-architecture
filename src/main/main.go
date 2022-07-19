package main


import (

"net/http"
"github.com/gorilla/mux"
"clean-go/src/main/factories/application/controllers"
)


func main() {
    r := mux.NewRouter()
    controllers.MakeUserHandlers(r)
    
    http.ListenAndServe(":8080", r)
}