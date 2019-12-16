package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MuxWeb() {
	r := mux.NewRouter()

	r.HandleFunc("/test/{aaa}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Println(vars["aaa"])
	})
	http.ListenAndServe(":80", r)
}
