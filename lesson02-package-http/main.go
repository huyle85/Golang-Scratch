package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demo_handleer)

	log.Println("Sever is starting..")
	err := http.ListenAndServe(":8080", nil) //localhoast 8080
	if err != nil{
		log.Fatal("Sever error: ", err)
	}
}

func demo_handleer(w http.ResponseWriter, r *http.Request){
	log.Println("%+v", r)
}
