package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
    Controllers "load-testing-golang/controllers"
    Init   "load-testing-golang/Init"  
)

func main() {
    router := mux.NewRouter().StrictSlash(true)

    Init.Init()
    

    router.HandleFunc("/Scenario2", Controllers.Scenario2).Methods("POST")

	log.Println("Starting API GOLANG PORT 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

