package main

import (
	"net/http"

	"github.com/jeypc/TUGAS/controller/roticontroller"
)

func main() {
	http.HandleFunc("/", roticontroller.Index)
	http.HandleFunc("/roti", roticontroller.Index)
	http.HandleFunc("/roti/index", roticontroller.Index)
	http.HandleFunc("/roti/add", roticontroller.Add)
	http.HandleFunc("/roti/edit", roticontroller.Edit)
	http.HandleFunc("/roti/delete", roticontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
