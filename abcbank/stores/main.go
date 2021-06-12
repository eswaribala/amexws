package stores

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectionHelper() *sql.DB {
	db, err := sql.Open("mysql", "root:vignesh@(127.0.0.1:3306)/abcbankdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTransaction(transactionId int64, amount int64, timestamp string, sender string, receiver string) (int64, error) {

	db := ConnectionHelper()
	queryString := "Insert into transaction (TransactionId,Amount,Time_Stamp,Sender,Receiver) values(?,?,?,?,?)"
	result, err := db.Exec(queryString, transactionId, amount, timestamp, sender, receiver)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	return result.RowsAffected()
}
