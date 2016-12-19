package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fid := vars["fid"]
	fmt.Fprint(w, fid)
}
