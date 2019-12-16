package web

import (
	"fmt"
	"net/http"
)

func WebStore() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "aaaa", r.URL.Path)
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
