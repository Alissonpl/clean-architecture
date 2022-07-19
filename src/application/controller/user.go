package controller

import (
"fmt"
"net/http"
"clean-go/src/domain/usecase/user"
)



func GetUser(service user.UseCase) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, service.GetUser())
	
})
}




