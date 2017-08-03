package routes

import {
    "Controller.Scenario1" AS  Scenario1
    "Controller.Scenario2" AS  Scenario2
}

func Route() {
	router.HandleFunc("/Scenario1", Scenario1).Methods('POST')
    router.HandleFunc("/Scenario2", Scenario2).Methods('POST')
}