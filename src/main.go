// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func MyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", ps.ByName("id"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/some/page/:id", MyHandler)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
