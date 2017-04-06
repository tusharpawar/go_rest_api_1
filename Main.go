package main 

import "log"
import "net/http"

func main(){
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081",router))
}