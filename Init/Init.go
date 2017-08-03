package Init

import (
	Model "load-testing-golang/models"
	Helper "load-testing-golang/Helper"  
)

func Init() {
	Mysql := Helper.Connection("mysql")
	Mysql.AutoMigrate(Model.Model{})
	defer Mysql.Close()
}