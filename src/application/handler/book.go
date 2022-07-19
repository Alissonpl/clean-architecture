package handler

import (
"fmt"
"github.com/gorilla/mux"
"clean-go/src/domain/usecase/book"
"net/http"
)



func getBook(service book.UseCase) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, service.GetBook())
	
})
}




func MakeBookHandlers(r *mux.Router, service book.UseCase) {
	r.Handle("/v1/book",
		getBook(service),
	).Methods("GET")

}
