package main

import (
	"net/http"
	"yujinGo/web3/myapplication"
)

func main(){
	http.ListenAndServe(":3000",myapplication.NewHttpHandler())
}