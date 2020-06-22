package main

import (
	"net/http"
	"yujinGo/web5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}