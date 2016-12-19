package api

import "net/http"

//执行检查
func Guard(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return f
}
