package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

const tableName string = "users"

type student struct {
	id         int
	name       string
	course     string
	created_at string
	updated_at string
}

func init() {
	db, err := sql.Open("mysql", "root:@/go_learn?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Connected....")
	dbm = db
}

// func connectDB() {
// 	db, err := sql.Open("mysql", "root:@/go_learn?parseTime=true")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("MySQL Connected....")
// 	dbm = db
// }

func main() {
	// var tableName string = "users"
	// connectDB()
	// createTable(tableName)
	insertData(tableName, "Satyajit Kumar", "PHP")
	insertData(tableName, "Satya Kr", "LARAVEL")
	insertData(tableName, "Akash Kumar", "GO")
	insertData(tableName, "Abhijit Kumar", "NODE JS")
	// updateData(tableName, "1", "Akash Kumar")
	getData(tableName)
	deleteData(tableName, "16")
	getData(tableName)
}

func chechTable(tablename string) bool {
	_, check := dbm.Query("SELECT * FROM " + tablename + ";")
	// fmt.Println(check)
	if check != nil {
		return false
	} else {
		return true
	}
}

func createTable(tableName string) {
	check := chechTable(tableName)
	if check {
		fmt.Println("Table " + tableName + " already exist")
	} else {
		query := `create table ` + tableName + `(
			id int(10) primary key auto_increment,
			name varchar(255) not null,
			course varchar(255) not null,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`
		_, err := dbm.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Users table created ...")
	}
}

func insertData(tableName, name, course string) {
	// name := "Satyajit Kumar"
	// course := "PHP"

	res, err := dbm.Exec(`INSERT INTO `+tableName+` (name, course) values (?, ?)`, name, course)
	if err != nil {
		log.Fatal(err)
	} else {
		lastId, _ := res.LastInsertId()
		fmt.Println("Insert record id is:", lastId)
	}
}

func updateData(tableName, id, name string) {
	_, err := dbm.Exec(`UPDATE `+tableName+` SET name=? WHERE id=?`, name, id)
	if err != nil {
		log.Fatal(err)
	} else {
		// count, _ := res.LastInsertId()
		// fmt.Println("Number of rows update:", count)
		fmt.Println("Record updated")
	}
}

func getData(tableName string) {
	rows, err := dbm.Query(`SELECT * FROM ` + tableName + `;`)
	if err != nil {
		log.Fatal(err)
	} else {
		var stu student
		for rows.Next() {
			rows.Scan(&stu.id, &stu.name, &stu.course, &stu.created_at, &stu.updated_at)
			fmt.Println(stu)
		}
	}
}

func deleteData(tableName, id string) {
	_, err := dbm.Exec(`DELETE FROM `+tableName+` WHERE id=?`, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record with id " + id + " deleted")
}
