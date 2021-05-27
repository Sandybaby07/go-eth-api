package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func TxHash(hash string, value string, gas string, gasPrice string, nonce string, data string, to string, pending bool) string {

	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/eth_transaction")
	fmt.Println(hash, value, gas, gasPrice, to, data, pending, "nonce")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "INSERT INTO transaction (Hash, Value, Gas, GasPrice, To_addr, Pending, Data, Nonce) VALUES (?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sql, hash, value, gas, gasPrice, to, pending, data, nonce)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Insert success.")
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

	return "ok"
}
