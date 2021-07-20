package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Ayush@1507"
	dbname   = "go"
)

func main() {
	

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	checkErr(err)
	defer db.Close()

	// fmt.Println("Successfully connected!")

	// -------------------------------------------------------------------------------------

	fmt.Println("# Inserting Row")

	// sql1:="INSERT INTO emp(id,name,age,email) values($1,$2,$3,$4)"
	// _,err=db.Exec(sql1,5,"Anurag Gupta",24,"anurag@gmail.com")

	sql1:=`INSERT INTO emp values($1,$2,$3,$4)`
	_,err=db.Exec(sql1,10,"Prakash",35,"prakash@gmail.com")

	checkErr(err)

	// -------------------------------------------------------------------------------------

	fmt.Println("# Updating Row")
	sql2 := `UPDATE emp SET name = $2, email = $3 WHERE id = $1;`
	_, err = db.Exec(sql2, 1, "Devesh", "devesh@gmail.com")
	checkErr(err)

	// -------------------------------------------------------------------------------------

	fmt.Println("# Delete Row")
	sql3 := `DELETE FROM emp WHERE id = $1;`
	_, err = db.Exec(sql3, 3)
	checkErr(err)

	// -------------------------------------------------------------------------------------

	fmt.Println("# Query Single Row")
	sql4 := `SELECT id, email FROM emp WHERE id=$1;`
	var email string
	var id int
	row := db.QueryRow(sql4, 5)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		checkErr(err)
	}
	
	// -------------------------------------------------------------------------------------

	fmt.Println("# Query Multiple Row")
	rows, err := db.Query("SELECT id, name FROM emp LIMIT $1", 3)
	checkErr(err)
  	defer rows.Close()
  	for rows.Next(){
		var id int
		var name string
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id, name)
	  }
  	err = rows.Err()
	checkErr(err)
}


func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

