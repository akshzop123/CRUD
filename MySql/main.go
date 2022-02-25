package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//emp struct

type emp struct {
	id                int
	name, email, role string
}

//open connection
func dbCon() *sql.DB {

	db, _:= sql.Open("mysql", "root:Akku@132@tcp(127.0.0.1:3306)/test")

	// if err != nil {
	// 	fmt.Println("error while opening")
	// }
	return db
}


var db = dbCon()
func main() {
	

	//insert
	Insert(db, emp{1, "emp1", "emp1@gmail.com", "Intern"})
	



	//select or getyid

	GetById(db, 1)


	//update

	UpdateById(db, "akshata", 1)

	//delete

	RemoveById(db, 3)


}

//Insert

func Insert(db *sql.DB, u emp) error {

	q := "INSERT INTO employee VALUES(?,?,?,?)"

	_, e := db.Exec(q, u.id, u.name, u.email, u.role)
	if e != nil {
		return e
	}
	//fmt.Println("A row inserted")
	return nil

}



func GetById(db *sql.DB, id int) (*emp, error) {

	var u emp

	q := "SELECT * FROM employee WHERE id=?"
	res := db.QueryRow(q, id)


	e := res.Scan(&u.id, &u.name, &u.email, &u.role)
	if e != nil {
		return nil, e
	}

	// fmt.Println(u.id,u.name,u.email,u.role)

	return &u, nil

}

//update

func UpdateById(db *sql.DB, name string, id int) error {

	q := "UPDATE employee SET name=? WHERE id=?"

	_, e := db.Exec(q, name, id)
	if e != nil {
		return e
	}

	return nil
}

//delete

func RemoveById(db *sql.DB, id int) error {

	q := "DELETE FROM employee WHERE id=?"

	_, e := db.Exec(q, id)
	if e != nil {
		return e
	}
	//defer del.Close()
	return nil
}