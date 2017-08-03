package Helper

import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func Default() *sql.DB {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	pass := "668482"
	dbname := "golang"

	dbinfo := "'"+user+"':"+pass+"@tcp("+host+":"+port+")/"+dbname+""
	
	db, err := sql.Open("mysql", dbinfo)
	 if(err != nil) {
			log.Println("Connection Error")
		}
	return db
}