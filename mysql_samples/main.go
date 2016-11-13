// main.go project main.go
package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	db, err := sql.Open("mysql", "root:@/bootadmin?charset=utf8")
	checkErr(err)
	defer db.Close()

	simpleSelectAll(db)
	singleFieldSelect(db)
	simpleInsertValue(db)

}

func simpleSelectAll(db *sql.DB) {
	rows, err := db.Query("select * from sys_user")
	checkErr(err)

	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		checkErr(err)

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}
		fmt.Println("--------------------------")
	}
}

func simpleInsertValue(db *sql.DB) {
	stmtIn, err := db.Prepare("insert into sys_user(name,account,passwd) values(?,?,?)")
	checkErr(err)
	defer stmtIn.Close()

	rs, err := stmtIn.Exec("kira", "1", "1")
	checkErr(err)

	id, err := rs.LastInsertId()
	checkErr(err)
	fmt.Println("Last  ID is ", id)
}
func singleFieldSelect(db *sql.DB) {
	stmtOut, err := db.Prepare("select account from sys_user where id=?")
	checkErr(err)
	defer stmtOut.Close()

	var account string
	stmtOut.QueryRow(1).Scan(&account)
	fmt.Println("Account is ", account)
}
