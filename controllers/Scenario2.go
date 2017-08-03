package controllers

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"runtime"
	Model "load-testing-golang/models"  
	Helper "load-testing-golang/Helper"
)

func Scenario2(w http.ResponseWriter, r *http.Request) {
	runtime.GOMAXPROCS(10)
	var channel = make(chan int)
	var total = 0
	var total_sum = func() {
        Mysql_Default := Helper.Default()
		defer Mysql_Default.Close()
		
		err := Mysql_Default.QueryRow("select SUM(amount) FROM models").Scan(&total)
		if err != nil {
		log.Println("error:", err)
		}
		channel<-total
    }

    go total_sum()

	Mysql := Helper.Connection("mysql")
	defer Mysql.Close()
		
	body, err := ioutil.ReadAll(r.Body)
		if(err != nil) {
			log.Println("Parse Error")
		}
	var data = Model.Model{}
	err = json.Unmarshal(body, &data)

	Mysql.Create(&Model.Model{
	    Amount          :data.Amount ,

	})
	var sum = <- channel
    b, err := json.Marshal(sum)
		if err != nil {
			log.Println("error:", err)
		}

	w.Write(b)

}

