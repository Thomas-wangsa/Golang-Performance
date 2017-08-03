package Helper

import (
	// "fmt"
	// "strings"
)

func MySQL() string {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	pass := "668482"
	dbname := "golang"


	dbinfo := "'"+user+"':"+pass+"@tcp("+host+":"+port+")/"+dbname+""

	return dbinfo
}