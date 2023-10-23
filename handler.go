package dumbo

import (
	"net/http"
)

type Handler http.HandlerFunc

type callback func(http.Server)


func methodNotAllowedHandler(methodsAllowed ...methodTyp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range methodsAllowed {
			w.Header().Add("Allow", reverseMethodMap[m])
		}
		w.WriteHeader(405)
		w.Write(nil)
	}
}


