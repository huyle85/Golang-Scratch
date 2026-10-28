package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demo_handleer)

	log.Println("Sever is starting..")
	err := http.ListenAndServe(":8080", nil) //localhoast 8080
	if err != nil {
		log.Fatal("Sever error: ", err)
	}
}

func demo_handleer(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	if r.Method != http.MethodGet {
		http.Error(w, "Phuong thuc 0 ho tro", http.StatusMethodNotAllowed)
		return
	}

	respone := map[string]string{
		"message": "Hello",
	}

	w.Header().Set("Content-Type", "application/json")

	// data, err := json.Marshal(respone)
	// if err != nil{
	// 	http.Error(w, "Loi", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(data)
	json.NewEncoder(w).Encode(respone)
}
