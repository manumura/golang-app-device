package main

import (
	"net/http"
)

// Application starts here.
import (
	"fmt"
	//"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe(":17172", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
